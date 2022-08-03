// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListOrgOKCode is the HTTP code returned for type ListOrgOK
const ListOrgOKCode int = 200

/*ListOrgOK OK

swagger:response listOrgOK
*/
type ListOrgOK struct {

	/*
	  In: Body
	*/
	Payload []*ListOrgOKBodyItems0 `json:"body,omitempty"`
}

// NewListOrgOK creates ListOrgOK with default headers values
func NewListOrgOK() *ListOrgOK {

	return &ListOrgOK{}
}

// WithPayload adds the payload to the list org o k response
func (o *ListOrgOK) WithPayload(payload []*ListOrgOKBodyItems0) *ListOrgOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list org o k response
func (o *ListOrgOK) SetPayload(payload []*ListOrgOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOrgOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListOrgOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListOrgUnauthorizedCode is the HTTP code returned for type ListOrgUnauthorized
const ListOrgUnauthorizedCode int = 401

/*ListOrgUnauthorized Unauthorized

swagger:response listOrgUnauthorized
*/
type ListOrgUnauthorized struct {
}

// NewListOrgUnauthorized creates ListOrgUnauthorized with default headers values
func NewListOrgUnauthorized() *ListOrgUnauthorized {

	return &ListOrgUnauthorized{}
}

// WriteResponse to the client
func (o *ListOrgUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListOrgForbiddenCode is the HTTP code returned for type ListOrgForbidden
const ListOrgForbiddenCode int = 403

/*ListOrgForbidden Forbidden

swagger:response listOrgForbidden
*/
type ListOrgForbidden struct {
}

// NewListOrgForbidden creates ListOrgForbidden with default headers values
func NewListOrgForbidden() *ListOrgForbidden {

	return &ListOrgForbidden{}
}

// WriteResponse to the client
func (o *ListOrgForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
