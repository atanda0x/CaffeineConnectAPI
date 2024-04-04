package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func ProductAPI(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Hey!!!!")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hey %s \n", d)
}
