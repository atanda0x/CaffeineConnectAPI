package handler

import (
	"net/http"

	"github.com/atanda0x/CaffeineConnectAPI/data"
)

// getProduct return the products from the datastore
func (p *Products) GetProducts(w http.ResponseWriter, _ *http.Request) {
	p.l.Println("Handle GET product")

	// fetch the product from datastore
	lp := data.GetProucts()

	// serialize the list to JSON
	err := lp.ToJOSN(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
