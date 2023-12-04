package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello Fuckers!")
    data, err := io.ReadAll(r.Body)
    if err != nil {
      http.Error(rw, "Oops, something went wrong here", http.StatusBadRequest)
      return
    }
    fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye FUCKERS!!!")
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		return
	}
}
