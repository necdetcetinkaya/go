package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err

	}
	p.ID = id
	productList[pos] = p
	return nil

}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil

		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1

}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "A",
		Description: "B",
		Price:       12.34,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "b",
		Description: "b",
		Price:       2.29,
		SKU:         "fjd34",
	},
	&Product{
		ID:          3,
		Name:        "c",
		Description: "c",
		Price:       2.39,
		SKU:         "TEST",
	},
}
