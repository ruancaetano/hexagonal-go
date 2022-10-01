package dto

import "github.com/ruancaetano/hexagonal-go/application"

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
}

func NewProductDto() *ProductDto {
	return &ProductDto{}
}

func (p *ProductDto) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Status = p.Status
	product.Price = p.Price

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
