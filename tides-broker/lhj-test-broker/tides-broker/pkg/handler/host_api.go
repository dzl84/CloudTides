package handler

import (
	//"crypto/tls"
	//"net/http"
	"log"
	"fmt"
	//"sync"
	//"sync/atomic"
	"os/exec"
	//"io/ioutil"
	"os"
	"strings"
	"strconv"
	//"float"

	"github.com/go-openapi/runtime/middleware"

	"tides-broker/pkg/config"
	"tides-broker/pkg/models"
	"tides-broker/pkg/restapi/operations/hosts_api"
)

// api for adding one host
func AddOneHost(params hosts_api.AddOneHostParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return NewAddOneHostUnauthorized()
	// }

	db := config.GetDB()

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.Body

	hostname := body.Hostname
	// search for hostname in the db and check if exists
	checkrst := db.Exec(fmt.Sprintf("SELECT * FROM hosts WHERE hostname = %s;", hostname))
	

	var rec = make(map[string]interface{})
    if checkrst.Find(rec); len(rec) > 0 {
        return hosts_api.NewAddOneHostBadRequest().WithPayload(&hosts_api.AddOneHostBadRequestBody{
			Message: "Hostname already exists",
		})
    }

	newHost := models.Host{
		Hostname: body.Hostname,
    	Policy: body.Policy,
    	Datacenter: body.Datacenter,
		Cluster: body.Cluster,
		IP: body.IP,
    	Port: body.Port,
		Sshkey: body.Sshkey,
    	Username: body.Username,
		Password: body.Password,
	}

	db.Create(&newHost)

	return hosts_api.NewAddOneHostCreated().WithPayload(&hosts_api.AddOneHostCreatedBody{
		Message: "success",
	})
}

// api for listing all the hosts
func GetHosts(params hosts_api.GetHostsParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return resource.NewListVsphereResourceUnauthorized()
	// }

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()

	// if VerifyAdmin(params.HTTPRequest) {
	// 	db.Find(&resources)
	// } else {
	// 	db.Where("user_id = ?", uid).Find(&resources)
	// }
	var hosts []*models.Host
	result := db.Find(&hosts)

	if result.Error  != nil {
		return hosts_api.NewGetHostsBadRequest().WithPayload(&hosts_api.GetHostsBadRequestBody{
			Message: "Cannot query the results",
		})
	}

	var response []*hosts_api.GetHostsOKBodyResultsItems0
	for _, tmpHost := range hosts {
		newResult := hosts_api.GetHostsOKBodyResultsItems0{
			Hostname: tmpHost.Hostname,
    	    Policy: tmpHost.Policy,
    	    Datacenter: tmpHost.Datacenter,
		    Cluster: tmpHost.Cluster,
		    IP: tmpHost.IP,
    	    Port: tmpHost.Port,
		    Sshkey: tmpHost.Sshkey,
		}

		response = append(response, &newResult)
	}

	return hosts_api.NewGetHostsOK().WithPayload(&hosts_api.GetHostsOKBody{
		Message: "success",
		Results: response,
	})
}


// api for deleting one host
func DeleteHost(params hosts_api.DeleteHostParams) middleware.Responder {
    // if !VerifyUser(params.HTTPRequest) {
    //  return NewAddOneHostUnauthorized()
    // }

    db := config.GetDB()

    // uid, _ := ParseUserIDFromToken(params.HTTPRequest)
    body := params.Body

    hostToDelete := body.Hostname

    rst := db.Exec(`DELETE FROM hosts WHERE hostname = $1;`, hostToDelete)
    if rst.Error != nil {
        return hosts_api.NewDeleteHostBadRequest().WithPayload(&hosts_api.DeleteHostBadRequestBody{
            Message: "Cannot delete host",
        })
    }

    return hosts_api.NewDeleteHostOK().WithPayload(&hosts_api.DeleteHostOKBody{
        Hostname: hostToDelete,
    })
}

// api for query hosts
func QueryHost(params hosts_api.QueryHostInfoParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return resource.NewListVsphereResourceUnauthorized()
	// }

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()

	// if VerifyAdmin(params.HTTPRequest) {
	// 	db.Find(&resources)
	// } else {
	// 	db.Where("user_id = ?", uid).Find(&resources)
	// }
	var hosts []*models.Host
	result := db.Find(&hosts)

	if result.Error  != nil {
		return hosts_api.NewQueryHostInfoBadRequest().WithPayload(&hosts_api.QueryHostInfoBadRequestBody{
			Message: "Cannot query the results",
		})
	}
	log.Print("Enter queryHosts handler")
	//result = make([]*hosts_api.QueryHostInfoOKBodyItems0, 0)
	// Query part
	// TODO:  Need to revise
	log.Print("Enter query.go")
	file, err := os.OpenFile("queryFile.ts", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Open file error")
	}
	port_address := "node-exporter:9100"
	str := `import { PrometheusDriver, QueryResult } from 'prometheus-query';

	const prom = new PrometheusDriver({
		endpoint: 'http://prometheus:9090/',
	});
	
	const port = "` + port_address + `"
	
	console.log(port)
	
	const cpu_query = 'avg by (instance) (rate(node_cpu_seconds_total{mode="idle", instance= "' + port + '"}[1m])) * 100';
	
	prom.instantQuery(cpu_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("LastQueryTime:", serie.value.time);
				console.log("percentCPU", serie.value.value);
			});
		})
		.catch(console.error);
	
	// memory
	const mem_free_query = 'node_memory_MemFree_bytes{instance= "' + port + '"}';
	prom.instantQuery(mem_free_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("currentRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	const mem_total_query = 'node_memory_MemTotal_bytes{instance= "' + port + '"}';
	prom.instantQuery(mem_total_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("totalRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	const mem_query = 'node_memory_MemFree_bytes{instance= "' + port + '"}/node_memory_MemTotal_bytes{instance= "' + port + '"} * 100';
	prom.instantQuery(mem_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("percentRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	// disk
	const disk_avail_query = 'node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_avail_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("currentDisk", serie.value.value);
			});
		})
		.catch(console.error);
	
	const disk_total_query = 'node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_total_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("totalDisk", serie.value.value);
			});
		})
		.catch(console.error);
	
	const disk_query = '(node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"} * 100) / node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("percentDisk", serie.value.value);
			});
		})
		.catch(console.error);`

	file.Write([]byte(str))
	defer file.Close()
	fmt.Println("Enter tsc")
	ts_cmd := exec.Command("tsc", "queryFile.ts")
	err = ts_cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		//return
	}

	js_cmd := exec.Command("node", "queryFile.js")

	stdout, err := js_cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	fmt.Print("Here is the output")
	fmt.Print(string(stdout))

	//res := string(stdout)
	res := make([]string, 0)
	res = append(res, string(stdout))

	// TODO:
	/*for _, tmpHost := range hosts{

	}*/
	var response []*hosts_api.QueryHostInfoOKBodyItems0
	for _, info := range res{
		hostInfo := hosts_api.QueryHostInfoOKBodyItems0{}
		dict := strings.Split(string(info), "\n")
		for _, i := range dict{
			if strings.Contains(i, ":") {
				temp := strings.Split(i, ":")
				if temp[0] == "LastQueryTime"{
					continue
				}
				hostInfo.Hostname = temp[0]
				hostInfo.Port = temp[1]
			}
			if strings.Contains(i, " "){
				temp := strings.Split(i, " ")
				var err error
				var t float64
				if temp[0] == "totalRAM"{
					t, err = strconv.ParseFloat(temp[1],32)
					hostInfo.TotalRAM = float32(t)
				}
				if temp[0] == "percentRAM"{
					t, err = strconv.ParseFloat(temp[1],32)
					hostInfo.PercentRAM = float32(t)
				}
				if temp[0] == "currentRAM"{
					t, err = strconv.ParseFloat(temp[1],32)
					hostInfo.CurrentRAM = float32(t)
				}
				if temp[0] == "percentCPU"{
					t, err = strconv.ParseFloat(temp[1],32)
					hostInfo.PercentCPU = float32(t)
				}
				if err!=nil{
					log.Print("error when parse response into struct")
				}
			}
		}
		response = append(response, &hostInfo)
	}

	return hosts_api.NewQueryHostInfoOK().WithPayload(response)
}

// api for query one host
func QueryOneHost(params hosts_api.QueryOneHostInfoParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return resource.NewListVsphereResourceUnauthorized()
	// }

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()

	// if VerifyAdmin(params.HTTPRequest) {
	// 	db.Find(&resources)
	// } else {
	// 	db.Where("user_id = ?", uid).Find(&resources)
	// }
	//var hosts []*models.Host
	//result := db.Find(&hosts)
	targetHost := *params.Hostname
	var hosts []*models.Host
	result := db.Find(&hosts, models.Host{Hostname: targetHost})

	if result.Error  != nil || len(hosts) != 1{
		return hosts_api.NewQueryOneHostInfoBadRequest().WithPayload(&hosts_api.QueryOneHostInfoBadRequestBody{
			Message: "Cannot query the results",
		})
	}
	log.Print("Enter queryHosts handler")

	//queryHost := hosts[0].Hostname
	portAdd := hosts[0].IP + ":" + strconv.FormatInt(hosts[0].Port,10)
	log.Print(portAdd)
	port_address := "node-exporter:9100"
	//port_address := portAdd

	//result = make([]*hosts_api.QueryHostInfoOKBodyItems0, 0)
	// Query part
	// TODO:  Need to revise
	log.Print("Enter query.go")
	file, err := os.OpenFile("queryFile.ts", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Open file error")
	}
	
	str := `import { PrometheusDriver, QueryResult } from 'prometheus-query';

	const prom = new PrometheusDriver({
		endpoint: 'http://prometheus:9090/',
	});
	
	const port = "` + port_address + `"
	
	console.log(port)
	
	const cpu_query = 'avg by (instance) (rate(node_cpu_seconds_total{mode="idle", instance= "' + port + '"}[1m])) * 100';
	
	prom.instantQuery(cpu_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("LastQueryTime:", serie.value.time);
				console.log("percentCPU", serie.value.value);
			});
		})
		.catch(console.error);
	
	// memory
	const mem_free_query = 'node_memory_MemFree_bytes{instance= "' + port + '"}';
	prom.instantQuery(mem_free_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("currentRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	const mem_total_query = 'node_memory_MemTotal_bytes{instance= "' + port + '"}';
	prom.instantQuery(mem_total_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("totalRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	const mem_query = 'node_memory_MemFree_bytes{instance= "' + port + '"}/node_memory_MemTotal_bytes{instance= "' + port + '"} * 100';
	prom.instantQuery(mem_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("percentRAM", serie.value.value);
			});
		})
		.catch(console.error);
	
	// disk
	const disk_avail_query = 'node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_avail_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("currentDisk", serie.value.value);
			});
		})
		.catch(console.error);
	
	const disk_total_query = 'node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_total_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("totalDisk", serie.value.value);
			});
		})
		.catch(console.error);
	
	const disk_query = '(node_filesystem_avail_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"} * 100) / node_filesystem_size_bytes{mountpoint="/",fstype!="rootfs",instance = "' + port + '"}';
	prom.instantQuery(disk_query)
		.then((res: QueryResult) => {
			const series = res.result;
			series.forEach((serie) => {
				console.log("percentDisk", serie.value.value);
			});
		})
		.catch(console.error);`

	file.Write([]byte(str))
	defer file.Close()
	fmt.Println("Enter tsc")
	ts_cmd := exec.Command("tsc", "queryFile.ts")
	err = ts_cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		//return
	}

	js_cmd := exec.Command("node", "queryFile.js")

	stdout, err := js_cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	fmt.Print("Here is the output")
	fmt.Print(string(stdout))

	//res := string(stdout)
	//res := make([]string, 0)
	//res = append(res, string(stdout))


	//var response hosts_api.QueryOneHostInfoOKBodyItems0
	hostInfo := hosts_api.QueryOneHostInfoOKBody{}
	dict := strings.Split(string(stdout), "\n")
	for _, i := range dict{
		if strings.Contains(i, ":") {
			temp := strings.Split(i, ":")
			if temp[0] == "LastQueryTime"{
				continue
			}
			hostInfo.Hostname = temp[0]
			hostInfo.Port = temp[1]
		}
		if strings.Contains(i, " "){
			temp := strings.Split(i, " ")
			var err error
			var t float64
			if temp[0] == "totalRAM"{
				t, err = strconv.ParseFloat(temp[1],32)
				hostInfo.TotalRAM = float32(t)
			}
			if temp[0] == "percentRAM"{
				t, err = strconv.ParseFloat(temp[1],32)
				hostInfo.PercentRAM = float32(t)
			}
			if temp[0] == "currentRAM"{
				t, err = strconv.ParseFloat(temp[1],32)
				hostInfo.CurrentRAM = float32(t)
			}
			if temp[0] == "percentCPU"{
				t, err = strconv.ParseFloat(temp[1],32)
				hostInfo.PercentCPU = float32(t)
			}
			if err!=nil{
				log.Print("error when parse response into struct")
			}
		}
	}	

	return hosts_api.NewQueryOneHostInfoOK().WithPayload(&hostInfo)
}