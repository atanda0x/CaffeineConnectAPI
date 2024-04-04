package main

import (
	"log"
	"net/http"
	"os"

	handler "github.com/atanda0x/CaffeineConnectAPI/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handler
	ph := handler.NewProducts(l)
	// Create a new serrve mux and register the handler
	sm := http.NewServeMux()
	sm.Handle("/", ph)

}
