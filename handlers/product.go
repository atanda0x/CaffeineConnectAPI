package handler

import (
	"log"
	"net/http"

	"github.com/atanda0x/CaffeineConnectAPI/data"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.handler interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request for a list of product
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// catch all
	// if no method is satisfied return err
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProduct return the products from the datastore
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET product")

	// fetch the produc t from datastore
	lp := data.GetProucts()

	// serialize the list to JSON
	err := lp.ToJOSN(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
