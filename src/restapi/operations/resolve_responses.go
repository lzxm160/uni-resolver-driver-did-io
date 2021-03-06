// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/iotexproject/uni-resolver-driver-did-io/src/models"
)

// ResolveOKCode is the HTTP code returned for type ResolveOK
const ResolveOKCode int = 200

/*ResolveOK OK

swagger:response resolveOK
*/
type ResolveOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ResolutionResult `json:"body,omitempty"`
}

// NewResolveOK creates ResolveOK with default headers values
func NewResolveOK() *ResolveOK {

	return &ResolveOK{}
}

// WithPayload adds the payload to the resolve o k response
func (o *ResolveOK) WithPayload(payload []*models.ResolutionResult) *ResolveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the resolve o k response
func (o *ResolveOK) SetPayload(payload []*models.ResolutionResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResolveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ResolutionResult, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ResolveBadRequestCode is the HTTP code returned for type ResolveBadRequest
const ResolveBadRequestCode int = 400

/*ResolveBadRequest invalid input!

swagger:response resolveBadRequest
*/
type ResolveBadRequest struct {
}

// NewResolveBadRequest creates ResolveBadRequest with default headers values
func NewResolveBadRequest() *ResolveBadRequest {

	return &ResolveBadRequest{}
}

// WriteResponse to the client
func (o *ResolveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ResolveInternalServerErrorCode is the HTTP code returned for type ResolveInternalServerError
const ResolveInternalServerErrorCode int = 500

/*ResolveInternalServerError error!

swagger:response resolveInternalServerError
*/
type ResolveInternalServerError struct {
}

// NewResolveInternalServerError creates ResolveInternalServerError with default headers values
func NewResolveInternalServerError() *ResolveInternalServerError {

	return &ResolveInternalServerError{}
}

// WriteResponse to the client
func (o *ResolveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
