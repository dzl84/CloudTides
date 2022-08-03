// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindTodosHandlerFunc turns a function with the right signature into a find todos handler
type FindTodosHandlerFunc func(FindTodosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindTodosHandlerFunc) Handle(params FindTodosParams) middleware.Responder {
	return fn(params)
}

// FindTodosHandler interface for that can handle valid find todos params
type FindTodosHandler interface {
	Handle(FindTodosParams) middleware.Responder
}

// NewFindTodos creates a new http.Handler for the find todos operation
func NewFindTodos(ctx *middleware.Context, handler FindTodosHandler) *FindTodos {
	return &FindTodos{Context: ctx, Handler: handler}
}

/* FindTodos swagger:route GET /addHost hosts findTodos

FindTodos find todos API

*/
type FindTodos struct {
	Context *middleware.Context
	Handler FindTodosHandler
}

func (o *FindTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindTodosParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
