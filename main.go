package main

import (
	"log"
	"net/http"
	"os"
	"time"

	handler "github.com/atanda0x/CaffeineConnectAPI/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handler
	ph := handler.NewProducts(l)

	// Create a new serrve mux and register the handler
	sm := mux.NewRouter()

	// Route
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.ServeHTTP)

	// create a new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.ListenAndServe()

}
