package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"helpdesk-ticketing-security/controller"
	"helpdesk-ticketing-security/service"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mvrilo/go-redoc"
)

func main() {
	// define routes

	router := mux.NewRouter()
	swagger := &redoc.Redoc{
		SpecPath: "/swagger.yaml",
		SpecFile: "swagger.yaml",
	}
	authController := &controller.AuthController{
		Service: service.AuthService{},
	}
	paswordRouter := router.PathPrefix("/password").Subrouter()
	passwordController := &controller.PasswordController{
		Service: service.PasswordService{},
	}

	paswordRouter.HandleFunc("/forgot", passwordController.Forgot).Methods("POST")
	paswordRouter.HandleFunc("/reset", passwordController.Reset).Methods("POST")
	router.HandleFunc("/login", authController.Login).Methods("POST")
	router.HandleFunc("/user", authController.User)
	router.HandleFunc("/token/refresh", authController.Refresh)
	router.PathPrefix("/swagger").Handler(swagger.Handler())
	router.HandleFunc("/health", controller.Ping)

	// launch server

	address := ":"

	if port, ok := os.LookupEnv("PORT"); ok && port != "" {
		address += port

	} else {
		address += "4000"
	}

	server := &http.Server{
		Addr:         address,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      router,
	}

	println("server starting on", address)
	log.Fatal(server.ListenAndServe())
}
