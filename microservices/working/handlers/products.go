package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"animo.com/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {

	return &Products{l}
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
		p.l.Println("PUT", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("got id", id)
		p.updateProducts(id, rw, r)

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle getProducts")
	lp := data.GetProducts()

	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
	//rw.Write(d)

}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle updateProducts")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)

	}
	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)

	}
	if err != nil {
		http.Error(rw, "Product can not be updated", http.StatusInternalServerError)
		return
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle addproduct")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)

	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)

}
