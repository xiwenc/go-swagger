package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/toolkit"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-2/restapi/operations"
	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-2/restapi/operations/todos"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = toolkit.JSONConsumer()

	api.JSONProducer = toolkit.JSONProducer()

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams) httpkit.Responder {
		return httpkit.NotImplemented("operation todos.AddOne has not yet been implemented")
	})
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams) httpkit.Responder {
		return httpkit.NotImplemented("operation todos.DestroyOne has not yet been implemented")
	})
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) httpkit.Responder {
		return httpkit.NotImplemented("operation todos.FindTodos has not yet been implemented")
	})
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams) httpkit.Responder {
		return httpkit.NotImplemented("operation todos.UpdateOne has not yet been implemented")
	})

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
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
