package handlers

import (
	"log"
	"net/http"

	"github.com/ricardoraposo/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// I'm not yet sold on this io.Writer thingy, I rather write this to the responseWriter myself
// with marshal and write methods
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

  rw.WriteHeader(http.StatusMethodNotAllowed)
}
