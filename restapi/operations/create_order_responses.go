// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"api.trafilea.com/models"
)

// CreateOrderOKCode is the HTTP code returned for type CreateOrderOK
const CreateOrderOKCode int = 200

/*
CreateOrderOK OK

swagger:response createOrderOK
*/
type CreateOrderOK struct {

	/*
	  In: Body
	*/
	Payload *models.OrderResponse `json:"body,omitempty"`
}

// NewCreateOrderOK creates CreateOrderOK with default headers values
func NewCreateOrderOK() *CreateOrderOK {

	return &CreateOrderOK{}
}

// WithPayload adds the payload to the create order o k response
func (o *CreateOrderOK) WithPayload(payload *models.OrderResponse) *CreateOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create order o k response
func (o *CreateOrderOK) SetPayload(payload *models.OrderResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateOrderBadRequestCode is the HTTP code returned for type CreateOrderBadRequest
const CreateOrderBadRequestCode int = 400

/*
CreateOrderBadRequest Bad request

swagger:response createOrderBadRequest
*/
type CreateOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateOrderBadRequest creates CreateOrderBadRequest with default headers values
func NewCreateOrderBadRequest() *CreateOrderBadRequest {

	return &CreateOrderBadRequest{}
}

// WithPayload adds the payload to the create order bad request response
func (o *CreateOrderBadRequest) WithPayload(payload string) *CreateOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create order bad request response
func (o *CreateOrderBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateOrderNotFoundCode is the HTTP code returned for type CreateOrderNotFound
const CreateOrderNotFoundCode int = 404

/*
CreateOrderNotFound Resource not found

swagger:response createOrderNotFound
*/
type CreateOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateOrderNotFound creates CreateOrderNotFound with default headers values
func NewCreateOrderNotFound() *CreateOrderNotFound {

	return &CreateOrderNotFound{}
}

// WithPayload adds the payload to the create order not found response
func (o *CreateOrderNotFound) WithPayload(payload string) *CreateOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create order not found response
func (o *CreateOrderNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateOrderInternalServerErrorCode is the HTTP code returned for type CreateOrderInternalServerError
const CreateOrderInternalServerErrorCode int = 500

/*
CreateOrderInternalServerError Internal server error

swagger:response createOrderInternalServerError
*/
type CreateOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateOrderInternalServerError creates CreateOrderInternalServerError with default headers values
func NewCreateOrderInternalServerError() *CreateOrderInternalServerError {

	return &CreateOrderInternalServerError{}
}

// WithPayload adds the payload to the create order internal server error response
func (o *CreateOrderInternalServerError) WithPayload(payload string) *CreateOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create order internal server error response
func (o *CreateOrderInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
