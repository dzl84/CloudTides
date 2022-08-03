// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemovePolicyHandlerFunc turns a function with the right signature into a remove policy handler
type RemovePolicyHandlerFunc func(RemovePolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemovePolicyHandlerFunc) Handle(params RemovePolicyParams) middleware.Responder {
	return fn(params)
}

// RemovePolicyHandler interface for that can handle valid remove policy params
type RemovePolicyHandler interface {
	Handle(RemovePolicyParams) middleware.Responder
}

// NewRemovePolicy creates a new http.Handler for the remove policy operation
func NewRemovePolicy(ctx *middleware.Context, handler RemovePolicyHandler) *RemovePolicy {
	return &RemovePolicy{Context: ctx, Handler: handler}
}

/* RemovePolicy swagger:route DELETE /policy/{id} policy removePolicy

remove a policy

*/
type RemovePolicy struct {
	Context *middleware.Context
	Handler RemovePolicyHandler
}

func (o *RemovePolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRemovePolicyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RemovePolicyNotFoundBody remove policy not found body
//
// swagger:model RemovePolicyNotFoundBody
type RemovePolicyNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this remove policy not found body
func (o *RemovePolicyNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove policy not found body based on context it is used
func (o *RemovePolicyNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemovePolicyNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemovePolicyNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RemovePolicyNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RemovePolicyOKBody remove policy o k body
//
// swagger:model RemovePolicyOKBody
type RemovePolicyOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this remove policy o k body
func (o *RemovePolicyOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove policy o k body based on context it is used
func (o *RemovePolicyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemovePolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemovePolicyOKBody) UnmarshalBinary(b []byte) error {
	var res RemovePolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
