// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPastUsageOKCode is the HTTP code returned for type GetPastUsageOK
const GetPastUsageOKCode int = 200

/*GetPastUsageOK success

swagger:response getPastUsageOK
*/
type GetPastUsageOK struct {

	/*
	  In: Body
	*/
	Payload []*GetPastUsageOKBodyItems0 `json:"body,omitempty"`
}

// NewGetPastUsageOK creates GetPastUsageOK with default headers values
func NewGetPastUsageOK() *GetPastUsageOK {

	return &GetPastUsageOK{}
}

// WithPayload adds the payload to the get past usage o k response
func (o *GetPastUsageOK) WithPayload(payload []*GetPastUsageOKBodyItems0) *GetPastUsageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get past usage o k response
func (o *GetPastUsageOK) SetPayload(payload []*GetPastUsageOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPastUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetPastUsageOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPastUsageUnauthorizedCode is the HTTP code returned for type GetPastUsageUnauthorized
const GetPastUsageUnauthorizedCode int = 401

/*GetPastUsageUnauthorized Unauthorized

swagger:response getPastUsageUnauthorized
*/
type GetPastUsageUnauthorized struct {
}

// NewGetPastUsageUnauthorized creates GetPastUsageUnauthorized with default headers values
func NewGetPastUsageUnauthorized() *GetPastUsageUnauthorized {

	return &GetPastUsageUnauthorized{}
}

// WriteResponse to the client
func (o *GetPastUsageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetPastUsageNotFoundCode is the HTTP code returned for type GetPastUsageNotFound
const GetPastUsageNotFoundCode int = 404

/*GetPastUsageNotFound resource not found

swagger:response getPastUsageNotFound
*/
type GetPastUsageNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetPastUsageNotFoundBody `json:"body,omitempty"`
}

// NewGetPastUsageNotFound creates GetPastUsageNotFound with default headers values
func NewGetPastUsageNotFound() *GetPastUsageNotFound {

	return &GetPastUsageNotFound{}
}

// WithPayload adds the payload to the get past usage not found response
func (o *GetPastUsageNotFound) WithPayload(payload *GetPastUsageNotFoundBody) *GetPastUsageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get past usage not found response
func (o *GetPastUsageNotFound) SetPayload(payload *GetPastUsageNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPastUsageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}