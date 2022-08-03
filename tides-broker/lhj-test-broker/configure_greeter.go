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

func addHost(item *models.Item) error {
	if item == nil {
		return errors.New(500, "host must be present")
	}

	hostsLock.Lock()
	defer hostsLock.Unlock()

	newID := newHostID()
	item.ID = newID
	hostsList[newID] = item

	return nil
}

func allHosts(since int64, limit int32) (result []*models.Item) {
	result = make([]*models.Item, 0)
	for id, item := range hostsList {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return result
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
		if err := addHost(params.Body); err != nil {
			return hosts.NewAddOneHostDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return hosts.NewAddOneHostCreated().WithPayload(params.Body)
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