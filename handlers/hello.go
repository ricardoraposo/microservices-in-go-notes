package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Fuckers!")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops, something went wrong here", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)
}
