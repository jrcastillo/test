// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"swagger-test/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger-test --name RedeServer --spec ../swagger.yml

func configureFlags(api *operations.RedeServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RedeServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.PutcapturaHandler == nil {
		api.PutcapturaHandler = operations.PutcapturaHandlerFunc(func(params operations.PutcapturaParams) middleware.Responder {
			return middleware.NotImplemented("operation .Putcaptura has not yet been implemented")
		})
	}
	if api.CapturaHandler == nil {
		api.CapturaHandler = operations.CapturaHandlerFunc(func(params operations.CapturaParams) middleware.Responder {
			return middleware.NotImplemented("operation .Captura has not yet been implemented")
		})
	}
	if api.RefundHandler == nil {
		api.RefundHandler = operations.RefundHandlerFunc(func(params operations.RefundParams) middleware.Responder {
			return middleware.NotImplemented("operation .Refund has not yet been implemented")
		})
	}
	if api.RefundlistHandler == nil {
		api.RefundlistHandler = operations.RefundlistHandlerFunc(func(params operations.RefundlistParams) middleware.Responder {
			return middleware.NotImplemented("operation .Refundlist has not yet been implemented")
		})
	}
	if api.RefundsTidHandler == nil {
		api.RefundsTidHandler = operations.RefundsTidHandlerFunc(func(params operations.RefundsTidParams) middleware.Responder {
			return middleware.NotImplemented("operation .RefundsTid has not yet been implemented")
		})
	}
	if api.TransaçãoHandler == nil {
		api.TransaçãoHandler = operations.TransaçãoHandlerFunc(func(params operations.TransaçãoParams) middleware.Responder {
			return middleware.NotImplemented("operation .Transação has not yet been implemented")
		})
	}
	if api.TransaçãoReferenceHandler == nil {
		api.TransaçãoReferenceHandler = operations.TransaçãoReferenceHandlerFunc(func(params operations.TransaçãoReferenceParams) middleware.Responder {
			return middleware.NotImplemented("operation .TransaçãoReference has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
