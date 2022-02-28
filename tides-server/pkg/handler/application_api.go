package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/application"
	"time"
)

var sshPool sync.Map
var chPool sync.Map
var mqPool sync.Map
var tokenIndex sync.Map
var (
	sshType    = "password" //password or key
	sshKeyPath = ""         //ssh id_rsa.id path"
)

var agentRfWebsshLink = map[string]string{
	"thinkcloud":  "http://www.cloudtides.org.cn:8000/webssh",
	"schoolcloud": "xxxxxx/webssh",
}

func AchieveHost(params application.AchieveHostParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}

	result := make([]*application.AchieveHostOKBodyItems0, 2)
	result[0] = &application.AchieveHostOKBodyItems0{
		AgentName: "thinkcloud",
	}
	result[1] = &application.AchieveHostOKBodyItems0{
		AgentName: "schoolcloud",
	}
	return application.NewAchieveHostOK().WithPayload(result)
}

func AchieveMQResult(params application.InstanceActionStatueParams) middleware.Responder {
	ch, has := chPool.Load(params.ReqBody.Token)
	if has {
		ch.(chan *OperateChan) <- &OperateChan{
			Combo:       params.ReqBody.Combo,
			SSHPort:     params.ReqBody.SSHPort,
			SSHUser:     params.ReqBody.SSHUser,
			SSHPassword: params.ReqBody.SSHPassword,
			SSHHost:     params.ReqBody.SSHHost,
		}
	}
	return application.NewInstanceActionStatueOK()
}

type Operate struct {
	Op      string
	Host    string
	User    string
	Pass    string
	Token   string
	Data    string
	Port    string
	CMD     string
	SshType string
}

type OperateChan struct {
	Combo       string
	SSHHost     string `json:"ssh_host"`
	SSHUser     string `json:"ssh_user"`
	SSHPassword string `json:"ssh_password"`
	SSHPort     string `json:"ssh_port"`
	Err         error
}

func CreateApplicationInstance(params application.CreateApplicationInstanceParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)

	body := params.ReqBody

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)

	var containerName string
	var containerPort string
	var containerToken string
	var cmd string

	switch body.AppType {
	case "jupyter":
		containerName = "jupyterlab-" + token
		containerPort = body.Port
		containerToken = token
		cmd = fmt.Sprintf("docker run -p %s:8888 -e JUPYTER_ENABLE_LAB=yes --rm -d --name %s jupyter/all-spark-notebook start-notebook.sh --NotebookApp.token='%s'", containerPort, containerName, containerToken)
	default:
		containerName = body.AppType + "-" + token
		cmd = fmt.Sprintf("docker run --hostname %s -it -v $HOME/data/%s:/data -w /data -d --name %s yinhaozheng/%s:latest", body.InstanceName, body.AppType, containerName, body.AppType)
	}

	//run remote shell
	cp := Operate{
		Op:      "Create",
		Token:   token,
		SshType: sshType,
		CMD:     cmd,
	}

	log.Print("Using MQ  body.AgentName: " + body.AgentName)

	msg, err := json.Marshal(cp)
	failOnError(err, "Failed to marshal json")

	ch := make(chan *OperateChan, 1)
	chPool.Store(token, ch)
	pushMsg2MQ(string(msg), body.AgentName, config.GetConfig().MqHost)
	combo, host, pwd, usr, port, err := readCh(ch, token)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewCreateApplicationInstanceForbidden().WithPayload(&application.CreateApplicationInstanceForbiddenBody{
			Msg: err.Error(),
		})
	}
	log.Println("cmd output:", string(combo))
	sshPortStr := port

	row := map[string]string{
		"agent":         body.AgentName,
		"containerID":   string(combo),
		"containerName": containerName,
		"instanceName":  body.InstanceName,
		"token":         token,
		"uid":           strconv.Itoa(int(uid)),
		"sshHost":       host,
		"sshUser":       usr,
		"sshPassword":   pwd,
		"sshType":       sshType,
		"sshKeyPath":    sshKeyPath,
		"sshPort":       sshPortStr,
		"appType":       body.AppType,
	}
	var extra interface{}
	switch body.AppType {
	case "jupyter":
		row["link"] = fmt.Sprintf("%s:%s/lab?token=%s", host, containerPort, token)
		row["port"] = containerPort
	default:
		row["link"] = config.GetConfig().WebSshServiceHost
		jsonMap := map[string]string{
			"sshuser":     row["sshUser"],
			"sshHost":     row["sshHost"],
			"sshPassword": row["sshPassword"],
			"sshPort":     row["sshPort"],
			"appType":     row["appType"],
			"cmd":         fmt.Sprintf("docker exec -it %s /bin/bash", row["containerID"][:12]),
		}

		if jsonByte, err := json.Marshal(&jsonMap); err != nil {
			log.Println(err)
		} else {
			jsonTemp := make(map[string]string)
			jsonTemp["base64"] = base64.StdEncoding.EncodeToString(jsonByte)
			row["extra"] = string(jsonByte)
			extra = jsonTemp
		}
	}
	var data models.Application
	if err := mapstructure.Decode(row, &data); err != nil {
		log.Println("map to struct error ", err)
	}
	db := config.GetDB()
	if err := db.Create(&data).Error; err != nil {
		log.Println("db install error", err)
	}
	log.Println(data)
	log.Println("row Id ", data.ID)
	tokenIndex.Store(token, data.ID)

	result := &application.CreateApplicationInstanceOKBody{
		//Link:  fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		Link:  row["link"],
		Token: token,
		Extra: extra,
	}
	return application.NewCreateApplicationInstanceOK().WithPayload(result)
}

func SearchApplicationInstance(params application.SearchInstanceParams) middleware.Responder {
	indexStock, has := tokenIndex.Load(params.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", params.Token).First(data)
	}

	msg := base64.StdEncoding.EncodeToString([]byte(data.Extra))

	return application.NewSearchInstanceOK().WithPayload(&application.SearchInstanceOKBody{
		Data: msg,
	})
}

func DeleteApplicationInstance(params application.DeleteApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}

	indexStock, has := tokenIndex.Load(params.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", params.Token).First(data)
	}

	cmd := fmt.Sprintf("docker kill %s", data.ContainerID[:12])
	log.Println("cmd ", cmd)
	//combo, err := execCmd(data.SshHost, port, data, cmd)

	//run remote shell
	cp := Operate{
		Op:      "Delete",
		Host:    data.SshHost,
		User:    data.SshUser,
		Pass:    data.SshPassword,
		Token:   data.Token,
		Port:    data.SshPort,
		SshType: sshType,
		CMD:     cmd,
	}

	var combo []byte

	msg, err := json.Marshal(cp)
	failOnError(err, "Failed to marshal json")

	log.Println("data.Agent : ", data.Agent)

	ch := make(chan *OperateChan, 1)
	chPool.Store(data.Token, ch)
	pushMsg2MQ(string(msg), data.Agent, config.GetConfig().MqHost)
	combo, _, _, _, _, err = readCh(ch, data.Token)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewDeleteApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	log.Println("index :", index)

	tokenIndex.Delete(params.Token)
	db.Delete(data)

	return application.NewDeleteApplicationInstanceOK().WithPayload(&application.DeleteApplicationInstanceOKBody{
		Msg: "success",
	})
}

func ListApplicationInstance(params application.ListApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var data []models.Application
	db.Where("uid = ?", fmt.Sprintf("%d", uid)).Find(&data)

	payload := make([]*application.ListApplicationInstanceOKBodyItems0, len(data), len(data))
	var jsonTemp map[string]string
	for i := range data {
		if len(data[i].Extra) > 0 {
			jsonTemp = make(map[string]string)
			jsonTemp["base64"] = base64.StdEncoding.EncodeToString([]byte(data[i].Extra))
		} else {
			jsonTemp = nil
		}

		tmp := &application.ListApplicationInstanceOKBodyItems0{
			Link:         data[i].Link,
			Token:        data[i].Token,
			InstanceName: data[i].InstanceName,
			AppType:      data[i].AppType,
			SSHHost:      data[i].SshHost,
			SSHPassword:  data[i].SshPassword,
			SSHPort:      data[i].SshPort,
			SSHUser:      data[i].SshUser,
			Port:         data[i].Port,
			AgentName:    data[i].Agent,
			Extra:        jsonTemp,
			CreateAt:     fmt.Sprintf("%d", data[i].CreatedAt.Unix()),
			RunningTime:  fmt.Sprintf("%d", time.Now().Unix()-data[i].CreatedAt.Unix()),
		}

		if value, has := agentRfWebsshLink[data[i].Agent]; has {
			tmp.Link = value
		}

		payload[i] = tmp
	}
	return application.NewListApplicationInstanceOK().WithPayload(payload)
}

func UpdateApplicationInstance(params application.UpdateApplicationInstanceParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	indexStock, has := tokenIndex.Load(body.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", body.Token).First(data)
	}

	cmd := fmt.Sprintf("docker kill %s", data.ContainerID[:12])
	log.Println("cmd ", cmd)
	port, err := strconv.Atoi(data.SshPort)
	combo, err := execCmd(data.SshHost, port, data, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewUpdateApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	log.Println("index :", index)

	tokenIndex.Delete(body.Token)
	db.Delete(data)

	myuuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Generating UUIDs error", err)
	}
	token := fmt.Sprintf("%x", sha256.Sum256(myuuid.Bytes()))
	log.Println("token", token)
	containerName := "jupyterlab-" + token
	containerPort := body.Port
	containerToken := token
	cmd = fmt.Sprintf("docker run -p %s:8888 -e JUPYTER_ENABLE_LAB=yes --rm -d --name %s jupyter/all-spark-notebook start-notebook.sh --NotebookApp.token='%s'", containerPort, containerName, containerToken)

	//run remote shell
	combo, err = execCmd(body.SSHHost, int(body.SSHPort), &models.Application{
		SshPassword: body.SSHPassword,
		SshUser:     body.SSHUser,
	}, cmd)

	if err != nil {
		log.Println("remote run cmd failed", cmd, err)
		return application.NewUpdateApplicationInstanceForbidden()
	}
	log.Println("cmd output:", string(combo))
	sshPortStr := strconv.Itoa(int(body.SSHPort))
	row := map[string]string{
		"link":          fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		"containerID":   string(combo),
		"containerName": containerName,
		"instanceName":  body.InstanceName,
		"token":         token,
		"uid":           strconv.Itoa(int(uid)),
		"sshHost":       body.SSHHost,
		"sshUser":       body.SSHUser,
		"sshPassword":   body.SSHPassword,
		"sshType":       sshType,
		"sshKeyPath":    sshKeyPath,
		"sshPort":       sshPortStr,
		"appType":       body.AppType,
		"port":          containerPort,
	}
	data = new(models.Application)
	if err := mapstructure.Decode(row, &data); err != nil {
		log.Println("map to struct error ", err)
	}
	if err := db.Create(&data).Error; err != nil {
		log.Println("db install error", err)
	}
	log.Println("row Id ", data.ID)
	tokenIndex.Store(token, data.ID)

	result := &application.UpdateApplicationInstanceOKBody{
		Link:  fmt.Sprintf("%s:%s/lab?token=%s", body.SSHHost, containerPort, token),
		Token: token,
	}
	return application.NewUpdateApplicationInstanceOK().WithPayload(result)

}

func WatchApplicationInstanceLogs(params application.WatchApplicationInstanceLogsParams) middleware.Responder {

	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}

	indexStock, has := tokenIndex.Load(params.Token)
	db := config.GetDB()
	var index uint
	data := new(models.Application)
	if has {
		index = indexStock.(uint)
		db.First(&data, index)
	} else {
		db.Where("token = ?", params.Token).First(data)
	}

	cmd := fmt.Sprintf("docker logs -n 200 %s", data.ContainerID[:12])
	port, _ := strconv.Atoi(data.SshPort)
	//run remote shell
	combo, err := execCmd(data.SshHost, port, data, cmd)
	if err != nil {
		log.Println("run remote shell error", err)
	}
	texts := strings.Split(string(combo), "\n")
	if len(texts) > 0 {
		texts = texts[:len(texts)-1]
	}
	payload := make([]*application.WatchApplicationInstanceLogsOKBodyItems0, len(texts), len(texts))
	for i, line := range texts {
		section := strings.Split(line, "] ")
		payload[i] = new(application.WatchApplicationInstanceLogsOKBodyItems0)
		if len(section) > 1 {
			payload[i].Content = section[1]
			row := strings.Split(section[0], " ")
			payload[i].Leve = strings.Replace(row[0], "[", "", 1)
			payload[i].Date = row[1]
			payload[i].Time = row[2]
			payload[i].Source = row[3]
		} else {
			payload[i].Content = section[0]
			payload[i].Leve = "None"
			payload[i].Date = "None"
			payload[i].Time = "None"
			payload[i].Source = "None"
		}

	}
	return application.NewWatchApplicationInstanceLogsOK().WithPayload(payload)
}

type WsHandler struct {
	Request *http.Request
}

func (ws WsHandler) WriteResponse(w http.ResponseWriter, r runtime.Producer) {

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(w, ws.Request, nil)
	defer conn.Close()
	if err != nil {
		log.Println("websocket error ", err)
	}

	go func() {
		for {
			code, text, err := conn.ReadMessage()
			log.Printf("code %d, text %s, err %s", code, string(text), err)
			if err != nil {
				return
			}
		}

	}()

	for {
		// todo get real logdata
		data := `
 {
        "content": "/usr/local/bin/start-notebook.sh: running hooks in /usr/local/bin/before-notebook.d",
        "date": "None",
        "leve": "None",
        "source": "None",
        "time": "None"
    }
`
		t := rand.Int63n(3)
		log.Println(t)
		time.Sleep(time.Second * time.Duration(t))

		err = conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("websocket write error ", err)
			return
		}
	}
}

func WsWatchApplicationInstanceLogs(params application.WsWatchApplicationInstanceLogsParams) middleware.Responder {
	wsStruct := &WsHandler{
		Request: params.HTTPRequest,
	}
	return wsStruct
}

func pushMsg2MQ(msg, name, dial string) {
	var producer sarama.AsyncProducer
	tmp, has := mqPool.Load(dial)
	if !has {
		producer = InitProducer(dial)
		log.Print("MQ host ", dial)

		mqPool.Store(dial, producer)
	} else {
		producer = tmp.(sarama.AsyncProducer)
	}

	ProducerSend(name, msg, producer)
	log.Printf(" Sent %s\n", msg)
}

func execCmd(sshHost string, sshPort int, data *models.Application, cmd string) ([]byte, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            data.SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshType == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(data.SshPassword)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(data.SshPassword)}
	}

	session := getSession(sshHost, sshPort, sshConfig)
	log.Println("cmd ", cmd)
	return session.CombinedOutput(cmd)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {

	key, err := ioutil.ReadFile(kPath)
	if err != nil {
		log.Println("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func readCh(ch chan *OperateChan, token string) (combos []byte, host, pwd, usr, port string, err error) {
	lastMsg := <-ch
	combo := lastMsg.Combo
	err = lastMsg.Err
	close(ch)
	chPool.Delete(token)
	log.Println(lastMsg)
	return []byte(combo), lastMsg.SSHHost, lastMsg.SSHPassword, lastMsg.SSHUser, lastMsg.SSHPort, err
}

func getSession(sshHost string, sshPort int, config *ssh.ClientConfig) (session *ssh.Session) {

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	log.Println("ssh addr ", addr)
	sshClientStock, has := sshPool.Load(addr)
	var sshClient *ssh.Client

	if !has {
		sshClient, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Println("created ssh client failed", err)
			if err := sshClient.Close(); err != nil {
				log.Println("close ssh client failed", err)
			}
		} else {
			log.Println("existing sshClient")
			sshPool.Store(addr, sshClient)
		}
	} else {
		log.Println("map existing sshClient")
		sshClient = sshClientStock.(*ssh.Client)
	}

	defer func(addr string) {
		err := recover()
		if err != nil {
			log.Println("recover ", err)
			sshClient, err := ssh.Dial("tcp", addr, config)
			log.Println("defer", err)
			log.Println("defer existing sshClient")
			sshPool.Delete(addr)
			sshPool.Store(addr, sshClient)
			session, err = sshClient.NewSession()
			log.Println("defer", err)
		}
	}(addr)

	//create ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Println("created session failed", err)
		if err := session.Close(); err != nil {
			log.Println("close session failed", err)
		}
	}

	return session
}

func UploadInstanceFile(params application.UploadInstanceFileParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	r := params.HTTPRequest
	token := params.Token
	log.Println("Instance token == ", token)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Get file Error", err)
	}
	defer file.Close()
	myconfig := config.GetConfig()
	//dirPath := fmt.Sprintf("%s/%d/%s", myconfig.TempStoragePath, uid, token)\
	dirPath := myconfig.TempStoragePath
	if err := createFile(dirPath); err != nil {
		log.Println("Create dir Error ", err)
	}
	filePath := fmt.Sprintf("%s/%s", dirPath, handler.Filename)
	log.Println(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("Create file Error", err)
	}
	defer f.Close()
	io.Copy(f, file)

	downLink := fmt.Sprintf("%s:%s/api/v1/application/instance/file/%d/%s/%s", myconfig.ServerIP, myconfig.ServerPort, uid, token, handler.Filename)
	log.Println(downLink)

	return application.NewUploadInstanceFileOK().WithPayload(&application.UploadInstanceFileOKBody{
		Downlink: downLink,
	})
}

func ListInstanceFiles(params application.ListInstanceFilesParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return application.NewCreateApplicationInstanceUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	token := params.Token
	myconfig := config.GetConfig()
	//dirPath := fmt.Sprintf("%s/%d/%s", myconfig.TempStoragePath, uid, token)
	dirPath := myconfig.TempStoragePath
	files, _ := ioutil.ReadDir(dirPath)
	payload := make([]*application.ListInstanceFilesOKBodyItems0, len(files))
	for i, f := range files {
		downLink := fmt.Sprintf("%s:%s/api/v1/application/instance/file/%d/%s/%s", "120.133.15.12", "8000", uid, token, f.Name())
		payload[i] = &application.ListInstanceFilesOKBodyItems0{
			Filename:   f.Name(),
			Filesize:   fmt.Sprintf("%d", f.Size()),
			Createtime: fmt.Sprintf("%d", f.ModTime().Unix()),
			Downlink:   downLink,
		}
		log.Println(f.ModTime())
	}
	return application.NewListInstanceFilesOK().WithPayload(payload)
}

type FileDown struct {
	FileName string
	FilePath string
}

func (down FileDown) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	file, _ := os.Open(down.FilePath)
	defer file.Close()
	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileStat, _ := file.Stat()
	rw.Header().Set("Content-Disposition", "attachment; filename="+down.FileName)
	rw.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	rw.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
	file.Seek(0, 0)
	io.Copy(rw, file)
}

func DownInstanceFile(params application.DownInstanceFileParams) middleware.Responder {
	//dirPath := fmt.Sprintf("%s/%s/%s", config.GetConfig().TempStoragePath, params.UID, params.Token)
	dirPath := config.GetConfig().TempStoragePath
	filename := fmt.Sprintf("%s/%s", dirPath, params.Name)
	log.Println(filename)
	st := new(FileDown)
	st.FilePath = filename
	st.FileName = params.Name
	return st
}
