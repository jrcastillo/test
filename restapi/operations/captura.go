// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CapturaHandlerFunc turns a function with the right signature into a captura handler
type CapturaHandlerFunc func(CapturaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CapturaHandlerFunc) Handle(params CapturaParams) middleware.Responder {
	return fn(params)
}

// CapturaHandler interface for that can handle valid captura params
type CapturaHandler interface {
	Handle(CapturaParams) middleware.Responder
}

// NewCaptura creates a new http.Handler for the captura operation
func NewCaptura(ctx *middleware.Context, handler CapturaHandler) *Captura {
	return &Captura{Context: ctx, Handler: handler}
}

/*Captura swagger:route GET /desenvolvedores/v1/transactions/{tid} captura

captura

Consulta da transação por TID.

*/
type Captura struct {
	Context *middleware.Context
	Handler CapturaHandler
}

func (o *Captura) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCapturaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
