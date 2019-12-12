// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RefundHandlerFunc turns a function with the right signature into a refund handler
type RefundHandlerFunc func(RefundParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RefundHandlerFunc) Handle(params RefundParams) middleware.Responder {
	return fn(params)
}

// RefundHandler interface for that can handle valid refund params
type RefundHandler interface {
	Handle(RefundParams) middleware.Responder
}

// NewRefund creates a new http.Handler for the refund operation
func NewRefund(ctx *middleware.Context, handler RefundHandler) *Refund {
	return &Refund{Context: ctx, Handler: handler}
}

/*Refund swagger:route POST /desenvolvedores/v1/transactions/{tid}/refunds refund

refund

Realiza o cancelamento da transação.

*/
type Refund struct {
	Context *middleware.Context
	Handler RefundHandler
}

func (o *Refund) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRefundParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
