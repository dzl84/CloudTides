// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceActionStatueHandlerFunc turns a function with the right signature into a instance action statue handler
type InstanceActionStatueHandlerFunc func(InstanceActionStatueParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InstanceActionStatueHandlerFunc) Handle(params InstanceActionStatueParams) middleware.Responder {
	return fn(params)
}

// InstanceActionStatueHandler interface for that can handle valid instance action statue params
type InstanceActionStatueHandler interface {
	Handle(InstanceActionStatueParams) middleware.Responder
}

// NewInstanceActionStatue creates a new http.Handler for the instance action statue operation
func NewInstanceActionStatue(ctx *middleware.Context, handler InstanceActionStatueHandler) *InstanceActionStatue {
	return &InstanceActionStatue{Context: ctx, Handler: handler}
}

/* InstanceActionStatue swagger:route POST /application/instance/action/statue application instanceActionStatue

instance action statue

*/
type InstanceActionStatue struct {
	Context *middleware.Context
	Handler InstanceActionStatueHandler
}

func (o *InstanceActionStatue) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewInstanceActionStatueParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// InstanceActionStatueBody instance action statue body
//
// swagger:model InstanceActionStatueBody
type InstanceActionStatueBody struct {

	// combo
	Combo string `json:"combo,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// ssh host
	SSHHost string `json:"ssh_host,omitempty"`

	// ssh password
	SSHPassword string `json:"ssh_password,omitempty"`

	// ssh port
	SSHPort string `json:"ssh_port,omitempty"`

	// ssh user
	SSHUser string `json:"ssh_user,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this instance action statue body
func (o *InstanceActionStatueBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance action statue body based on context it is used
func (o *InstanceActionStatueBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstanceActionStatueBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstanceActionStatueBody) UnmarshalBinary(b []byte) error {
	var res InstanceActionStatueBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// InstanceActionStatueForbiddenBody instance action statue forbidden body
//
// swagger:model InstanceActionStatueForbiddenBody
type InstanceActionStatueForbiddenBody struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this instance action statue forbidden body
func (o *InstanceActionStatueForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance action statue forbidden body based on context it is used
func (o *InstanceActionStatueForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstanceActionStatueForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstanceActionStatueForbiddenBody) UnmarshalBinary(b []byte) error {
	var res InstanceActionStatueForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
