package main

import (
	"log"
	"net/http"

	"github.com/huangx8/oapi-codegen-share/example/api"
	"github.com/huangx8/oapi-codegen-share/example/pkg/server"
)

func main() {
	mux := http.NewServeMux()
	cardCenter := server.NewServer()
	handler := api.HandlerFromMux(cardCenter, mux)
	httpServer := &http.Server{
		Handler: handler,
		Addr:    ":8888",
	}

	log.Fatal(httpServer.ListenAndServe())
}
