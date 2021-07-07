// Code generated by go-swagger; DO NOT EDIT.

package vmtemp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateVMTempHandlerFunc turns a function with the right signature into a update VM temp handler
type UpdateVMTempHandlerFunc func(UpdateVMTempParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVMTempHandlerFunc) Handle(params UpdateVMTempParams) middleware.Responder {
	return fn(params)
}

// UpdateVMTempHandler interface for that can handle valid update VM temp params
type UpdateVMTempHandler interface {
	Handle(UpdateVMTempParams) middleware.Responder
}

// NewUpdateVMTemp creates a new http.Handler for the update VM temp operation
func NewUpdateVMTemp(ctx *middleware.Context, handler UpdateVMTempHandler) *UpdateVMTemp {
	return &UpdateVMTemp{Context: ctx, Handler: handler}
}

/* UpdateVMTemp swagger:route PUT /vmtemp vmtemp updateVmTemp

update VMTemp

*/
type UpdateVMTemp struct {
	Context *middleware.Context
	Handler UpdateVMTempHandler
}

func (o *UpdateVMTemp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateVMTempParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateVMTempBody update VM temp body
//
// swagger:model UpdateVMTempBody
type UpdateVMTempBody struct {

	// disk
	Disk int64 `json:"disk,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// ports
	Ports string `json:"ports,omitempty"`

	// vcpu
	Vcpu int64 `json:"vcpu,omitempty"`

	// vmem
	Vmem int64 `json:"vmem,omitempty"`
}

// Validate validates this update VM temp body
func (o *UpdateVMTempBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update VM temp body based on context it is used
func (o *UpdateVMTempBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVMTempBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVMTempBody) UnmarshalBinary(b []byte) error {
	var res UpdateVMTempBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateVMTempNotFoundBody update VM temp not found body
//
// swagger:model UpdateVMTempNotFoundBody
type UpdateVMTempNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update VM temp not found body
func (o *UpdateVMTempNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update VM temp not found body based on context it is used
func (o *UpdateVMTempNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVMTempNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVMTempNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateVMTempNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateVMTempOKBody update VM temp o k body
//
// swagger:model UpdateVMTempOKBody
type UpdateVMTempOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update VM temp o k body
func (o *UpdateVMTempOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update VM temp o k body based on context it is used
func (o *UpdateVMTempOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateVMTempOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateVMTempOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateVMTempOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
