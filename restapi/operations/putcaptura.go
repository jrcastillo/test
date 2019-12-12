// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutcapturaHandlerFunc turns a function with the right signature into a putcaptura handler
type PutcapturaHandlerFunc func(PutcapturaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutcapturaHandlerFunc) Handle(params PutcapturaParams) middleware.Responder {
	return fn(params)
}

// PutcapturaHandler interface for that can handle valid putcaptura params
type PutcapturaHandler interface {
	Handle(PutcapturaParams) middleware.Responder
}

// NewPutcaptura creates a new http.Handler for the putcaptura operation
func NewPutcaptura(ctx *middleware.Context, handler PutcapturaHandler) *Putcaptura {
	return &Putcaptura{Context: ctx, Handler: handler}
}

/*Putcaptura swagger:route PUT /desenvolvedores/v1/transactions/{tid} putcaptura

captura

Confirma a autorização da transação (captura).

*/
type Putcaptura struct {
	Context *middleware.Context
	Handler PutcapturaHandler
}

func (o *Putcaptura) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutcapturaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}