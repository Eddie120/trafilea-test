// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"api.trafilea.com/models"
)

// ModifyQuantityOKCode is the HTTP code returned for type ModifyQuantityOK
const ModifyQuantityOKCode int = 200

/*
ModifyQuantityOK OK

swagger:response modifyQuantityOK
*/
type ModifyQuantityOK struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewModifyQuantityOK creates ModifyQuantityOK with default headers values
func NewModifyQuantityOK() *ModifyQuantityOK {

	return &ModifyQuantityOK{}
}

// WithPayload adds the payload to the modify quantity o k response
func (o *ModifyQuantityOK) WithPayload(payload *models.Product) *ModifyQuantityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify quantity o k response
func (o *ModifyQuantityOK) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyQuantityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ModifyQuantityBadRequestCode is the HTTP code returned for type ModifyQuantityBadRequest
const ModifyQuantityBadRequestCode int = 400

/*
ModifyQuantityBadRequest Bad request

swagger:response modifyQuantityBadRequest
*/
type ModifyQuantityBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewModifyQuantityBadRequest creates ModifyQuantityBadRequest with default headers values
func NewModifyQuantityBadRequest() *ModifyQuantityBadRequest {

	return &ModifyQuantityBadRequest{}
}

// WithPayload adds the payload to the modify quantity bad request response
func (o *ModifyQuantityBadRequest) WithPayload(payload string) *ModifyQuantityBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify quantity bad request response
func (o *ModifyQuantityBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyQuantityBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ModifyQuantityNotFoundCode is the HTTP code returned for type ModifyQuantityNotFound
const ModifyQuantityNotFoundCode int = 404

/*
ModifyQuantityNotFound Resource not found

swagger:response modifyQuantityNotFound
*/
type ModifyQuantityNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewModifyQuantityNotFound creates ModifyQuantityNotFound with default headers values
func NewModifyQuantityNotFound() *ModifyQuantityNotFound {

	return &ModifyQuantityNotFound{}
}

// WithPayload adds the payload to the modify quantity not found response
func (o *ModifyQuantityNotFound) WithPayload(payload string) *ModifyQuantityNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify quantity not found response
func (o *ModifyQuantityNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyQuantityNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ModifyQuantityInternalServerErrorCode is the HTTP code returned for type ModifyQuantityInternalServerError
const ModifyQuantityInternalServerErrorCode int = 500

/*
ModifyQuantityInternalServerError Internal server error

swagger:response modifyQuantityInternalServerError
*/
type ModifyQuantityInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewModifyQuantityInternalServerError creates ModifyQuantityInternalServerError with default headers values
func NewModifyQuantityInternalServerError() *ModifyQuantityInternalServerError {

	return &ModifyQuantityInternalServerError{}
}

// WithPayload adds the payload to the modify quantity internal server error response
func (o *ModifyQuantityInternalServerError) WithPayload(payload string) *ModifyQuantityInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the modify quantity internal server error response
func (o *ModifyQuantityInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ModifyQuantityInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
