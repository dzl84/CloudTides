// Code generated by go-swagger; DO NOT EDIT.

package hosts_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteHostHandlerFunc turns a function with the right signature into a delete host handler
type DeleteHostHandlerFunc func(DeleteHostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHostHandlerFunc) Handle(params DeleteHostParams) middleware.Responder {
	return fn(params)
}

// DeleteHostHandler interface for that can handle valid delete host params
type DeleteHostHandler interface {
	Handle(DeleteHostParams) middleware.Responder
}

// NewDeleteHost creates a new http.Handler for the delete host operation
func NewDeleteHost(ctx *middleware.Context, handler DeleteHostHandler) *DeleteHost {
	return &DeleteHost{Context: ctx, Handler: handler}
}

/* DeleteHost swagger:route POST /deleteHost hosts_api deleteHost

delete a host

*/
type DeleteHost struct {
	Context *middleware.Context
	Handler DeleteHostHandler
}

func (o *DeleteHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteHostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteHostBadRequestBody delete host bad request body
//
// swagger:model DeleteHostBadRequestBody
type DeleteHostBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete host bad request body
func (o *DeleteHostBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete host bad request body based on context it is used
func (o *DeleteHostBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteHostBody use hostname to delete host
//
// swagger:model DeleteHostBody
type DeleteHostBody struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this delete host body
func (o *DeleteHostBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete host body based on context it is used
func (o *DeleteHostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteHostOKBody delete host o k body
//
// swagger:model DeleteHostOKBody
type DeleteHostOKBody struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this delete host o k body
func (o *DeleteHostOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete host o k body based on context it is used
func (o *DeleteHostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
