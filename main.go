package main

import (
	"log"
	"net/http"
	"os"

	"github.com/atanda0x/CaffeineConnectAPI/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.ProductAPI(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
