// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"api.trafilea.com/controllers"
	"crypto/tls"
	"github.com/rs/cors"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"api.trafilea.com/restapi/operations"
)

//go:generate swagger generate server --target ../../trafilea-test --name Trafilea --spec ../swagger.json --principal interface{}

func configureFlags(api *operations.TrafileaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TrafileaAPI) http.Handler {
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	container := buildContainer()

	invoke := func(initializer interface{}) {
		err := container.Invoke(initializer)
		if err != nil {
			panic(err)
		}
	}

	invoke(func(controller *controllers.ShoppingController) {
		api.AddNewProductHandler = operations.AddNewProductHandlerFunc(func(params operations.AddNewProductParams) middleware.Responder {
			return controller.AddProductToCart(params)
		})
	})

	invoke(func(controller *controllers.ShoppingController) {
		api.CreateCartHandler = operations.CreateCartHandlerFunc(func(params operations.CreateCartParams) middleware.Responder {
			return controller.CreateCart(params)
		})
	})

	invoke(func(controller *controllers.ShoppingController) {
		api.CreateOrderHandler = operations.CreateOrderHandlerFunc(func(params operations.CreateOrderParams) middleware.Responder {
			return controller.CreateOrder(params)
		})
	})

	invoke(func(controller *controllers.ShoppingController) {
		api.ModifyQuantityHandler = operations.ModifyQuantityHandlerFunc(func(params operations.ModifyQuantityParams) middleware.Responder {
			return controller.UpdateQuantity(params)
		})
	})

	api.PreServerShutdown = func() {}

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
	return cors.AllowAll().Handler(handler)
}
