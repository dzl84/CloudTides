// Code generated by go-swagger; DO NOT EDIT.

package vmtemp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteVMTempOKCode is the HTTP code returned for type DeleteVMTempOK
const DeleteVMTempOKCode int = 200

/*DeleteVMTempOK deletion success

swagger:response deleteVmTempOK
*/
type DeleteVMTempOK struct {

	/*
	  In: Body
	*/
	Payload *DeleteVMTempOKBody `json:"body,omitempty"`
}

// NewDeleteVMTempOK creates DeleteVMTempOK with default headers values
func NewDeleteVMTempOK() *DeleteVMTempOK {

	return &DeleteVMTempOK{}
}

// WithPayload adds the payload to the delete Vm temp o k response
func (o *DeleteVMTempOK) WithPayload(payload *DeleteVMTempOKBody) *DeleteVMTempOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Vm temp o k response
func (o *DeleteVMTempOK) SetPayload(payload *DeleteVMTempOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMTempOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVMTempUnauthorizedCode is the HTTP code returned for type DeleteVMTempUnauthorized
const DeleteVMTempUnauthorizedCode int = 401

/*DeleteVMTempUnauthorized Unauthorized

swagger:response deleteVmTempUnauthorized
*/
type DeleteVMTempUnauthorized struct {
}

// NewDeleteVMTempUnauthorized creates DeleteVMTempUnauthorized with default headers values
func NewDeleteVMTempUnauthorized() *DeleteVMTempUnauthorized {

	return &DeleteVMTempUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteVMTempUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteVMTempForbiddenCode is the HTTP code returned for type DeleteVMTempForbidden
const DeleteVMTempForbiddenCode int = 403

/*DeleteVMTempForbidden Forbidden

swagger:response deleteVmTempForbidden
*/
type DeleteVMTempForbidden struct {
}

// NewDeleteVMTempForbidden creates DeleteVMTempForbidden with default headers values
func NewDeleteVMTempForbidden() *DeleteVMTempForbidden {

	return &DeleteVMTempForbidden{}
}

// WriteResponse to the client
func (o *DeleteVMTempForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteVMTempNotFoundCode is the HTTP code returned for type DeleteVMTempNotFound
const DeleteVMTempNotFoundCode int = 404

/*DeleteVMTempNotFound resource not found

swagger:response deleteVmTempNotFound
*/
type DeleteVMTempNotFound struct {

	/*
	  In: Body
	*/
	Payload *DeleteVMTempNotFoundBody `json:"body,omitempty"`
}

// NewDeleteVMTempNotFound creates DeleteVMTempNotFound with default headers values
func NewDeleteVMTempNotFound() *DeleteVMTempNotFound {

	return &DeleteVMTempNotFound{}
}

// WithPayload adds the payload to the delete Vm temp not found response
func (o *DeleteVMTempNotFound) WithPayload(payload *DeleteVMTempNotFoundBody) *DeleteVMTempNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Vm temp not found response
func (o *DeleteVMTempNotFound) SetPayload(payload *DeleteVMTempNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMTempNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
