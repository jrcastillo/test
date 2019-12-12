// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RefundlistHandlerFunc turns a function with the right signature into a refundlist handler
type RefundlistHandlerFunc func(RefundlistParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RefundlistHandlerFunc) Handle(params RefundlistParams) middleware.Responder {
	return fn(params)
}

// RefundlistHandler interface for that can handle valid refundlist params
type RefundlistHandler interface {
	Handle(RefundlistParams) middleware.Responder
}

// NewRefundlist creates a new http.Handler for the refundlist operation
func NewRefundlist(ctx *middleware.Context, handler RefundlistHandler) *Refundlist {
	return &Refundlist{Context: ctx, Handler: handler}
}

/*Refundlist swagger:route GET /desenvolvedores/v1/transactions/{tid}/refunds/{refundId} refundlist

refund list

Consulta de cancelamento por refundId.

*/
type Refundlist struct {
	Context *middleware.Context
	Handler RefundlistHandler
}

func (o *Refundlist) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRefundlistParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}