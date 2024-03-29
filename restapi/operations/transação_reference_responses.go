// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// TransaçãoReferenceOKCode is the HTTP code returned for type TransaçãoReferenceOK
const TransaçãoReferenceOKCode int = 200

/*TransaçãoReferenceOK OK

swagger:response transaçãoReferenceOK
*/
type TransaçãoReferenceOK struct {
}

// NewTransaçãoReferenceOK creates TransaçãoReferenceOK with default headers values
func NewTransaçãoReferenceOK() *TransaçãoReferenceOK {

	return &TransaçãoReferenceOK{}
}

// WriteResponse to the client
func (o *TransaçãoReferenceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
