package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mvrilo/go-redoc"
)

func main() {
	// define routes

	router := mux.NewRouter()
	swagger := &redoc.Redoc{
		SpecPath: "/swagger.yaml",
		SpecFile: "swagger.yaml",
	}

	router.PathPrefix("/swagger").Handler(swagger.Handler())

	// launch server

	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      router,
	}

	println("server starting on :8000")
	server.ListenAndServe()
}
