package handler

import (
	"net/http"
	"strconv"

	"github.com/atanda0x/CaffeineConnectAPI/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
	}
	p.l.Println("Handle POST Products", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProducts(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}
