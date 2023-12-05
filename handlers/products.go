package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	// fetches the products from the database (not actual database)
	lp := data.GetProducts()

	// serializes the list to JSON
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

    err = data.UpdateProduct(id, prod)
    if err == data.ErrProductNotFound {
        http.Error(rw, "Product not found", http.StatusNotFound)
        return
    }

    if err != nil {
        http.Error(rw, "Product not found", http.StatusInternalServerError)
        return
    }
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		re := regexp.MustCompile(`/([0-9]+)`)
		g := re.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invali URL, more than one ID", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invali URL, more than one capture group", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Unable to convert string to number", http.StatusBadRequest)
			return
		}

        p.updateProduct(id, rw, r)
        return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
