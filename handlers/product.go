package handler

import (
	"log"
	"net/http"

	"github.com/atanda0x/CaffeineConnectAPI/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProucts()
	err := lp.ToJOSN(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
