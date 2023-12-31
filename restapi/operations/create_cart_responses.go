// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"api.trafilea.com/models"
)

// CreateCartOKCode is the HTTP code returned for type CreateCartOK
const CreateCartOKCode int = 200

/*
CreateCartOK UUID for the new cart

swagger:response createCartOK
*/
type CreateCartOK struct {

	/*
	  In: Body
	*/
	Payload *models.NewCartResponse `json:"body,omitempty"`
}

// NewCreateCartOK creates CreateCartOK with default headers values
func NewCreateCartOK() *CreateCartOK {

	return &CreateCartOK{}
}

// WithPayload adds the payload to the create cart o k response
func (o *CreateCartOK) WithPayload(payload *models.NewCartResponse) *CreateCartOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cart o k response
func (o *CreateCartOK) SetPayload(payload *models.NewCartResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCartOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCartBadRequestCode is the HTTP code returned for type CreateCartBadRequest
const CreateCartBadRequestCode int = 400

/*
CreateCartBadRequest Bad request

swagger:response createCartBadRequest
*/
type CreateCartBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateCartBadRequest creates CreateCartBadRequest with default headers values
func NewCreateCartBadRequest() *CreateCartBadRequest {

	return &CreateCartBadRequest{}
}

// WithPayload adds the payload to the create cart bad request response
func (o *CreateCartBadRequest) WithPayload(payload string) *CreateCartBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cart bad request response
func (o *CreateCartBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCartBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
