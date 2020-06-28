// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"tides-server/pkg/models"
)

// ResourceInfoHandlerFunc turns a function with the right signature into a resource info handler
type ResourceInfoHandlerFunc func(ResourceInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ResourceInfoHandlerFunc) Handle(params ResourceInfoParams) middleware.Responder {
	return fn(params)
}

// ResourceInfoHandler interface for that can handle valid resource info params
type ResourceInfoHandler interface {
	Handle(ResourceInfoParams) middleware.Responder
}

// NewResourceInfo creates a new http.Handler for the resource info operation
func NewResourceInfo(ctx *middleware.Context, handler ResourceInfoHandler) *ResourceInfo {
	return &ResourceInfo{Context: ctx, Handler: handler}
}

/*ResourceInfo swagger:route GET /resource/get_details resource resourceInfo

returns detailed info of resources belonging to a user

*/
type ResourceInfo struct {
	Context *middleware.Context
	Handler ResourceInfoHandler
}

func (o *ResourceInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewResourceInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ResourceInfoNotFoundBody resource info not found body
//
// swagger:model ResourceInfoNotFoundBody
type ResourceInfoNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this resource info not found body
func (o *ResourceInfoNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ResourceInfoNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResourceInfoNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ResourceInfoNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ResourceInfoOKBody resource info o k body
//
// swagger:model ResourceInfoOKBody
type ResourceInfoOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// results
	Results []*models.ResourceInfoItem `json:"results"`
}

// Validate validates this resource info o k body
func (o *ResourceInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ResourceInfoOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceInfoOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ResourceInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResourceInfoOKBody) UnmarshalBinary(b []byte) error {
	var res ResourceInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}