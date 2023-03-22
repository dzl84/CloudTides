package main

import (
	"flag"
	//"fmt"
	//"log"
	//"os"
	//"path/filepath"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-openapi/loads"
	//"github.com/go-openapi/runtime/middleware"
	//"github.com/go-openapi/swag"

	"tides-broker/pkg/restapi"
	"tides-broker/pkg/restapi/operations"
	// "tides-broker/pkg/models"
)

var portFlag = flag.Int("port", 2999, "Port to run this service on")

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var sshPool sync.Map


type Operate struct {
	Op      string
	Token   string
	Data    string
	CMD     string
	SshType string

	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}

type Host struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
}

var url string
var mqHost string
var mqTopic string

var consumer sarama.Consumer

// Callback func
type ConsumerCallback func(data []byte)

// Init Consumer
func InitConsumer(hosts string) {
	
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	failOnError(err, "unable to create kafka client")

	consumer, err = sarama.NewConsumerFromClient(client)
	failOnError(err, "NewAsyncProducerFromClient Err")
	log.Print("lalala for try")

}

func LoopConsumer(topic string, callback ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	failOnError(err, "unable to create partitionConsumer")
	defer partitionConsumer.Close()

	for {
		msg := <-partitionConsumer.Messages()
		if callback != nil {
			callback(msg.Value)
		}
	}
}

func Close() {
	if consumer != nil {
		consumer.Close()
	}
}

func TopicCallBack(msg []byte) {
	log.Printf("Received a message: %s", msg)
	op := &Operate{}
	err := json.Unmarshal(msg, op)
	failOnError(err, "Failed convert JSON string to struct")
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		log.Println(rand.Intn(len(hosts)))
	}
	host := hosts[rand.Intn(len(hosts))]
	var combo []byte
	if op.Op == "delete" {
		combo, err = execCmd(op.Host, op.User, op.Pass, op.SshType, op.Port, op.CMD)
	} else {
		combo, err = execCmd(host.Host, host.User, host.Pass, op.SshType, host.Port, op.CMD)

	}
	failOnError(err, "Failed to connect host by ssh")
	log.Printf("Received a combo: %s", combo)

	//SSHHost     string
	//SSHUser     string
	//SSHPassword string
	//SSHPort     string
	var values map[string]string
	if op.Op == "delete" {
		values = map[string]string{"combo": string(combo),
			"token":        op.Token,
			"ssh_host":     op.Host,
			"ssh_user":     op.User,
			"ssh_password": op.Pass,
			"ssh_port":     op.Port,
		}
	} else {
		values = map[string]string{"combo": string(combo),
			"token":        op.Token,
			"ssh_host":     host.Host,
			"ssh_user":     host.User,
			"ssh_password": host.Pass,
			"ssh_port":     host.Port,
		}
	}

	if err != nil {
		values["error"] = err.Error()
	} else {
		values["error"] = ""
	}

	log.Printf("Received a values: %v", values)

	jsonData, err := json.Marshal(values)
	failOnError(err, "Failed to convent json")
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	failOnError(err, "Failed to connect host")
	log.Printf("Received a resp: %s", resp.Status)

}

var hosts []*Host

func init() {
	mqHost = os.Getenv("MQ_HOST")
	mqTopic = os.Getenv("AGENT_NAME")
	url = os.Getenv("SERVER_URL")
	mqTopic="thinkcloud"
	mqHost="120.133.15.12:9092"
	url="http://backend:8033/api/v1/application/instance/action/statue"
	InitConsumer(mqHost)
	log.Print(os.Args[0])
	yamlFile, err := ioutil.ReadFile("hosts.yaml")
	if err != nil{
		log.Print("read file hosts.yaml wrong")
	}
	log.Print(string(yamlFile))
	log.Print(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &hosts)
	failOnError(err, "Failed to Unmarshal")
	log.Print(hosts[0])
	log.Print(hosts[0])
}



func execCmd(sshHost, SshUser, SshPassword, sshType, sshPort, cmd string) ([]byte, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if sshType == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(SshPassword)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(SshPassword)}
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

func getSession(sshHost string, sshPort string, config *ssh.ClientConfig) (session *ssh.Session) {

	addr := fmt.Sprintf("%s:%s", sshHost, sshPort)
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


func main() {
	log.Print("main start")
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewBrokerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// GetGreetingHandler greets the given name,
	// in case the name is not given, it will default to World
	/*api.GetGreetingHandler = operations.GetGreetingHandlerFunc(
		func(params operations.GetGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}

			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})
	*/
	// serve API
	server.ConfigureAPI()
	go func(){
		if err := server.Serve(); err != nil {
			log.Fatalln(err)
		}
	}()
	
	log.Print("MQ host tryyy", mqHost)

	forever := make(chan bool)

	LoopConsumer(mqTopic, TopicCallBack)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	Close()
	
}