// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ValidateResourceOKCode is the HTTP code returned for type ValidateResourceOK
const ValidateResourceOKCode int = 200

/*ValidateResourceOK returns the list of data centers belonging to the host

swagger:response validateResourceOK
*/
type ValidateResourceOK struct {

	/*
	  In: Body
	*/
	Payload *ValidateResourceOKBody `json:"body,omitempty"`
}

// NewValidateResourceOK creates ValidateResourceOK with default headers values
func NewValidateResourceOK() *ValidateResourceOK {

	return &ValidateResourceOK{}
}

// WithPayload adds the payload to the validate resource o k response
func (o *ValidateResourceOK) WithPayload(payload *ValidateResourceOKBody) *ValidateResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate resource o k response
func (o *ValidateResourceOK) SetPayload(payload *ValidateResourceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ValidateResourceUnauthorizedCode is the HTTP code returned for type ValidateResourceUnauthorized
const ValidateResourceUnauthorizedCode int = 401

/*ValidateResourceUnauthorized Unauthorized

swagger:response validateResourceUnauthorized
*/
type ValidateResourceUnauthorized struct {
}

// NewValidateResourceUnauthorized creates ValidateResourceUnauthorized with default headers values
func NewValidateResourceUnauthorized() *ValidateResourceUnauthorized {

	return &ValidateResourceUnauthorized{}
}

// WriteResponse to the client
func (o *ValidateResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ValidateResourceNotFoundCode is the HTTP code returned for type ValidateResourceNotFound
const ValidateResourceNotFoundCode int = 404

/*ValidateResourceNotFound resource not found

swagger:response validateResourceNotFound
*/
type ValidateResourceNotFound struct {

	/*
	  In: Body
	*/
	Payload *ValidateResourceNotFoundBody `json:"body,omitempty"`
}

// NewValidateResourceNotFound creates ValidateResourceNotFound with default headers values
func NewValidateResourceNotFound() *ValidateResourceNotFound {

	return &ValidateResourceNotFound{}
}

// WithPayload adds the payload to the validate resource not found response
func (o *ValidateResourceNotFound) WithPayload(payload *ValidateResourceNotFoundBody) *ValidateResourceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate resource not found response
func (o *ValidateResourceNotFound) SetPayload(payload *ValidateResourceNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
