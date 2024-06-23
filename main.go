package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/huangx8/oapi-codegen-share/example/api"
	"github.com/huangx8/oapi-codegen-share/example/pkg/server"
)

func someMiddleware(handler http.Handler) http.Handler {
	log.Println("some middleware " + uuid.NewString())
	return handler
}

func main() {
	mux := http.NewServeMux()
	cardCenter := server.NewServer()
	options := api.StdHTTPServerOptions{
		BaseRouter:  mux,
		Middlewares: []api.MiddlewareFunc{someMiddleware},
	}
	handler := api.HandlerWithOptions(cardCenter, options)
	httpServer := &http.Server{
		Handler: handler,
		Addr:    ":8888",
	}

	log.Fatal(httpServer.ListenAndServe())
}
