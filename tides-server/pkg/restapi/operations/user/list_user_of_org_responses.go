// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListUserOfOrgOKCode is the HTTP code returned for type ListUserOfOrgOK
const ListUserOfOrgOKCode int = 200

/*ListUserOfOrgOK OK

swagger:response listUserOfOrgOK
*/
type ListUserOfOrgOK struct {

	/*
	  In: Body
	*/
	Payload []*ListUserOfOrgOKBodyItems0 `json:"body,omitempty"`
}

// NewListUserOfOrgOK creates ListUserOfOrgOK with default headers values
func NewListUserOfOrgOK() *ListUserOfOrgOK {

	return &ListUserOfOrgOK{}
}

// WithPayload adds the payload to the list user of org o k response
func (o *ListUserOfOrgOK) WithPayload(payload []*ListUserOfOrgOKBodyItems0) *ListUserOfOrgOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user of org o k response
func (o *ListUserOfOrgOK) SetPayload(payload []*ListUserOfOrgOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserOfOrgOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListUserOfOrgOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListUserOfOrgUnauthorizedCode is the HTTP code returned for type ListUserOfOrgUnauthorized
const ListUserOfOrgUnauthorizedCode int = 401

/*ListUserOfOrgUnauthorized Unauthorized

swagger:response listUserOfOrgUnauthorized
*/
type ListUserOfOrgUnauthorized struct {
}

// NewListUserOfOrgUnauthorized creates ListUserOfOrgUnauthorized with default headers values
func NewListUserOfOrgUnauthorized() *ListUserOfOrgUnauthorized {

	return &ListUserOfOrgUnauthorized{}
}

// WriteResponse to the client
func (o *ListUserOfOrgUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListUserOfOrgForbiddenCode is the HTTP code returned for type ListUserOfOrgForbidden
const ListUserOfOrgForbiddenCode int = 403

/*ListUserOfOrgForbidden Forbidden

swagger:response listUserOfOrgForbidden
*/
type ListUserOfOrgForbidden struct {
}

// NewListUserOfOrgForbidden creates ListUserOfOrgForbidden with default headers values
func NewListUserOfOrgForbidden() *ListUserOfOrgForbidden {

	return &ListUserOfOrgForbidden{}
}

// WriteResponse to the client
func (o *ListUserOfOrgForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
