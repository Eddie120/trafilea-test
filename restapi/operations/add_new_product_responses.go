// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"api.trafilea.com/models"
)

// AddNewProductOKCode is the HTTP code returned for type AddNewProductOK
const AddNewProductOKCode int = 200

/*
AddNewProductOK OK

swagger:response addNewProductOK
*/
type AddNewProductOK struct {

	/*
	  In: Body
	*/
	Payload models.ListProducts `json:"body,omitempty"`
}

// NewAddNewProductOK creates AddNewProductOK with default headers values
func NewAddNewProductOK() *AddNewProductOK {

	return &AddNewProductOK{}
}

// WithPayload adds the payload to the add new product o k response
func (o *AddNewProductOK) WithPayload(payload models.ListProducts) *AddNewProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new product o k response
func (o *AddNewProductOK) SetPayload(payload models.ListProducts) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ListProducts{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddNewProductBadRequestCode is the HTTP code returned for type AddNewProductBadRequest
const AddNewProductBadRequestCode int = 400

/*
AddNewProductBadRequest Bad request

swagger:response addNewProductBadRequest
*/
type AddNewProductBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddNewProductBadRequest creates AddNewProductBadRequest with default headers values
func NewAddNewProductBadRequest() *AddNewProductBadRequest {

	return &AddNewProductBadRequest{}
}

// WithPayload adds the payload to the add new product bad request response
func (o *AddNewProductBadRequest) WithPayload(payload string) *AddNewProductBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new product bad request response
func (o *AddNewProductBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewProductBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddNewProductNotFoundCode is the HTTP code returned for type AddNewProductNotFound
const AddNewProductNotFoundCode int = 404

/*
AddNewProductNotFound Resource not found

swagger:response addNewProductNotFound
*/
type AddNewProductNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddNewProductNotFound creates AddNewProductNotFound with default headers values
func NewAddNewProductNotFound() *AddNewProductNotFound {

	return &AddNewProductNotFound{}
}

// WithPayload adds the payload to the add new product not found response
func (o *AddNewProductNotFound) WithPayload(payload string) *AddNewProductNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new product not found response
func (o *AddNewProductNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewProductNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddNewProductInternalServerErrorCode is the HTTP code returned for type AddNewProductInternalServerError
const AddNewProductInternalServerErrorCode int = 500

/*
AddNewProductInternalServerError Internal server error

swagger:response addNewProductInternalServerError
*/
type AddNewProductInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddNewProductInternalServerError creates AddNewProductInternalServerError with default headers values
func NewAddNewProductInternalServerError() *AddNewProductInternalServerError {

	return &AddNewProductInternalServerError{}
}

// WithPayload adds the payload to the add new product internal server error response
func (o *AddNewProductInternalServerError) WithPayload(payload string) *AddNewProductInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new product internal server error response
func (o *AddNewProductInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewProductInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
