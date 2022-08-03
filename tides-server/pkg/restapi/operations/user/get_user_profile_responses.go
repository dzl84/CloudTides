// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUserProfileOKCode is the HTTP code returned for type GetUserProfileOK
const GetUserProfileOKCode int = 200

/*GetUserProfileOK returns user profile

swagger:response getUserProfileOK
*/
type GetUserProfileOK struct {

	/*
	  In: Body
	*/
	Payload *GetUserProfileOKBody `json:"body,omitempty"`
}

// NewGetUserProfileOK creates GetUserProfileOK with default headers values
func NewGetUserProfileOK() *GetUserProfileOK {

	return &GetUserProfileOK{}
}

// WithPayload adds the payload to the get user profile o k response
func (o *GetUserProfileOK) WithPayload(payload *GetUserProfileOKBody) *GetUserProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user profile o k response
func (o *GetUserProfileOK) SetPayload(payload *GetUserProfileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserProfileUnauthorizedCode is the HTTP code returned for type GetUserProfileUnauthorized
const GetUserProfileUnauthorizedCode int = 401

/*GetUserProfileUnauthorized Unauthorized

swagger:response getUserProfileUnauthorized
*/
type GetUserProfileUnauthorized struct {
}

// NewGetUserProfileUnauthorized creates GetUserProfileUnauthorized with default headers values
func NewGetUserProfileUnauthorized() *GetUserProfileUnauthorized {

	return &GetUserProfileUnauthorized{}
}

// WriteResponse to the client
func (o *GetUserProfileUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetUserProfileNotFoundCode is the HTTP code returned for type GetUserProfileNotFound
const GetUserProfileNotFoundCode int = 404

/*GetUserProfileNotFound resource not found

swagger:response getUserProfileNotFound
*/
type GetUserProfileNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetUserProfileNotFoundBody `json:"body,omitempty"`
}

// NewGetUserProfileNotFound creates GetUserProfileNotFound with default headers values
func NewGetUserProfileNotFound() *GetUserProfileNotFound {

	return &GetUserProfileNotFound{}
}

// WithPayload adds the payload to the get user profile not found response
func (o *GetUserProfileNotFound) WithPayload(payload *GetUserProfileNotFoundBody) *GetUserProfileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user profile not found response
func (o *GetUserProfileNotFound) SetPayload(payload *GetUserProfileNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProfileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
