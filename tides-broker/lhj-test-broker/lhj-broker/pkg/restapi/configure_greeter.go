// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"log"
	"fmt"
	"sync"
	"sync/atomic"
	"os/exec"
	"io/ioutil"
	"os"
	
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	
	
	"lhj-broker/pkg/restapi/operations"
	"lhj-broker/pkg/restapi/operations/hosts"
	"lhj-broker/pkg/models"
)
//go:generate swagger generate server --target ../../gen --name Greeter --spec ../../swagger/swagger.yml --principal interface{} --exclude-main

var hostsList = make(map[int64]*models.Item)
var lastID int64

var hostsLock = &sync.Mutex{}

func newHostID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

//func addHost(item *models.Item) error {
func addHost(item hosts.AddOneHostBody) (hosts.AddOneHostBody, error) {
	//if item == nil {
	//	return errors.New(500, "host must be present")
	//}

	hostsLock.Lock()
	defer hostsLock.Unlock()

	newID := newHostID()
	item.ID = newID
	newItem := models.Item{
		ID: 		item.ID,
		IP:			item.IP,
		Completed:	item.Completed,
	}
	hostsList[newID] = &newItem

	return item, nil
}

//func allHosts(since int64, limit int32) (result []*models.Item) {
func allHosts(since int64, limit int32) (result []*hosts.FindHostsOKBodyItems0) {
	//result = make([]*models.Item, 0)
	result = make([]*hosts.FindHostsOKBodyItems0, 0)
	for id, item := range hostsList {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			returnItem := hosts.FindHostsOKBodyItems0{
				ID: 		item.ID,
				IP:			item.IP,
				Completed:	item.Completed,
			}
			//item.ID = 
			result = append(result, &returnItem)
		}
	}
	return result
}

func queryHosts(since int64, limit int32)(result []*hosts.QueryHostInfoOKBodyItems0, err error){
	log.Print("Enter queryHosts handler")
	result = make([]*hosts.QueryHostInfoOKBodyItems0, 0)
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
	log.Print(os.Getwd())
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	
	for _, f := range files {
		log.Print(f.Name())
	}
	fmt.Println("Enter tsc")
	ts_cmd := exec.Command("tsc", "queryFile.ts")
	err = ts_cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		return
	}

	js_cmd := exec.Command("node", "queryFile.js")

	stdout, err := js_cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		log.Print(err.Error())
		return
	}
	fmt.Print("Here is the output")
	fmt.Print(string(stdout))

	res := string(stdout)

	for id, item := range hostsList {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			returnItem := hosts.QueryHostInfoOKBodyItems0{
				ID: 		item.ID,
				IP:			item.IP,
				Info:		res,
			}
			//item.ID = 
			result = append(result, &returnItem)
		}
	}
	fmt.Print(result)
	return result, nil
}


func configureFlags(api *operations.GreeterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GreeterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	log.Print("Enter configureAPI")
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.TxtProducer = runtime.TextProducer()
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(func(params operations.GetGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}
	
			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})
	
	api.HostsFindHostsHandler = hosts.FindHostsHandlerFunc(func(params hosts.FindHostsParams) middleware.Responder {
		mergedParams := hosts.NewFindHostsParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return hosts.NewFindHostsOK().WithPayload(allHosts(*mergedParams.Since, *mergedParams.Limit))
	})

	api.HostsAddOneHostHandler = hosts.AddOneHostHandlerFunc(
		func(params hosts.AddOneHostParams) middleware.Responder {
		//newItem := hosts.AddOneHostBody{}

		item, err := addHost(params.Body)
		newItem := hosts.AddOneHostCreatedBody{
				ID: 		item.ID,
				IP:			item.IP,
				Completed:	item.Completed,
		}
		if err != nil{
			return hosts.NewAddOneHostDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		/*if _, err := addHost(params.Body); err != nil {
			return hosts.NewAddOneHostDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}*/
		return hosts.NewAddOneHostCreated().WithPayload(&newItem)
		})
	
	api.HostsQueryHostInfoHandler = hosts.QueryHostInfoHandlerFunc(
		func(params hosts.QueryHostInfoParams) middleware.Responder{
			mergedParams := hosts.NewQueryHostInfoParams()
			log.Print("Enter host query")
			mergedParams.Since = swag.Int64(0)
			if params.Since != nil {
				mergedParams.Since = params.Since
			}
			if params.Limit != nil {
				mergedParams.Limit = params.Limit
			}
			result, err := queryHosts(*mergedParams.Since, *mergedParams.Limit)
			if err != nil {
				return hosts.NewQueryHostInfoDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
			}
			return hosts.NewQueryHostInfoOK().WithPayload(result)
		})
	
	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}