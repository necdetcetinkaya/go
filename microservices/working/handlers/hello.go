package handlers

import (
	"io/ioutil"
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

	h.l.Println("Hello World tt")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "error Occured", http.StatusBadRequest)

		return
	}
	log.Printf("Data %s\n", d)

}
