package handler

import (
	"encoding/json"
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

func (p *Products) ServeHttp(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProuct()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	w.Write(d)
}
