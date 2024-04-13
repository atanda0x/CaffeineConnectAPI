// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/atanda0x/CaffeineConnectAPI/data"
)

// A list of products return in the response
// swagger:response
type productsResponse struct {
	// All product in the system
	//in: body
	Body []data.Product
}

// swagger:response
type noContent struct{}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	//in: path
	//required: true
	ID int `json:"id"`
}

// keyProduct type
type KeyProduct struct{}

// Product is a http.handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Middleware validate the product in the request and calls next
func (p *Products) MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarsal json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Add the Product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain
		next.ServeHTTP(w, req)
	})
}
