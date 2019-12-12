// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RefundNoContentCode is the HTTP code returned for type RefundNoContent
const RefundNoContentCode int = 204

/*RefundNoContent No Content

swagger:response refundNoContent
*/
type RefundNoContent struct {
}

// NewRefundNoContent creates RefundNoContent with default headers values
func NewRefundNoContent() *RefundNoContent {

	return &RefundNoContent{}
}

// WriteResponse to the client
func (o *RefundNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}
