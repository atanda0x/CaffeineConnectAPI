package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handler "github.com/atanda0x/CaffeineConnectAPI/handlers"
	"github.com/go-openapi/runtime/middleware"
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
	getRouter.HandleFunc("/", ph.GetProducts)
	getRouter.Use(ph.MiddlewareProductValidator)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidator)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidator)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create a new server
	s := http.Server{
		Addr:         ":9090",           // Address
		Handler:      sm,                // Set default handler
		ErrorLog:     l,                 // Set the logger for the server
		ReadTimeout:  5 * time.Second,   // Max time to  read request from the client
		WriteTimeout: 10 * time.Second,  // Max time to write response to the client
		IdleTimeout:  120 * time.Second, // Max time for connection using TCP keep-alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
