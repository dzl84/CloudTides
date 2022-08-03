// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateProjectHandlerFunc turns a function with the right signature into a update project handler
type UpdateProjectHandlerFunc func(UpdateProjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateProjectHandlerFunc) Handle(params UpdateProjectParams) middleware.Responder {
	return fn(params)
}

// UpdateProjectHandler interface for that can handle valid update project params
type UpdateProjectHandler interface {
	Handle(UpdateProjectParams) middleware.Responder
}

// NewUpdateProject creates a new http.Handler for the update project operation
func NewUpdateProject(ctx *middleware.Context, handler UpdateProjectHandler) *UpdateProject {
	return &UpdateProject{Context: ctx, Handler: handler}
}

/* UpdateProject swagger:route PUT /project/{id} project updateProject

update boinc project

*/
type UpdateProject struct {
	Context *middleware.Context
	Handler UpdateProjectHandler
}

func (o *UpdateProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateProjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateProjectBody update project body
//
// swagger:model UpdateProjectBody
type UpdateProjectBody struct {

	// has account manager
	HasAccountManager bool `json:"hasAccountManager,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this update project body
func (o *UpdateProjectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update project body based on context it is used
func (o *UpdateProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateProjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateProjectBody) UnmarshalBinary(b []byte) error {
	var res UpdateProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateProjectOKBody update project o k body
//
// swagger:model UpdateProjectOKBody
type UpdateProjectOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update project o k body
func (o *UpdateProjectOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update project o k body based on context it is used
func (o *UpdateProjectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateProjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateProjectOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateProjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
