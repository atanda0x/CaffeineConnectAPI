package handler

import (
	"net/http"

	"github.com/atanda0x/CaffeineConnectAPI/data"
)

// AddProduct add new Product to the list of product in the datastore
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

}
