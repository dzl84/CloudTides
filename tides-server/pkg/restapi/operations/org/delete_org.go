// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteOrgHandlerFunc turns a function with the right signature into a delete org handler
type DeleteOrgHandlerFunc func(DeleteOrgParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOrgHandlerFunc) Handle(params DeleteOrgParams) middleware.Responder {
	return fn(params)
}

// DeleteOrgHandler interface for that can handle valid delete org params
type DeleteOrgHandler interface {
	Handle(DeleteOrgParams) middleware.Responder
}

// NewDeleteOrg creates a new http.Handler for the delete org operation
func NewDeleteOrg(ctx *middleware.Context, handler DeleteOrgHandler) *DeleteOrg {
	return &DeleteOrg{Context: ctx, Handler: handler}
}

/* DeleteOrg swagger:route DELETE /org/{id} org deleteOrg

delete Org

*/
type DeleteOrg struct {
	Context *middleware.Context
	Handler DeleteOrgHandler
}

func (o *DeleteOrg) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteOrgParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteOrgNotFoundBody delete org not found body
//
// swagger:model DeleteOrgNotFoundBody
type DeleteOrgNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete org not found body
func (o *DeleteOrgNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete org not found body based on context it is used
func (o *DeleteOrgNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOrgNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOrgNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteOrgNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteOrgOKBody delete org o k body
//
// swagger:model DeleteOrgOKBody
type DeleteOrgOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete org o k body
func (o *DeleteOrgOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete org o k body based on context it is used
func (o *DeleteOrgOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOrgOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOrgOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteOrgOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}