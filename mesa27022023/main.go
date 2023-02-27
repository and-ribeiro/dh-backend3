package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save(product Product) {
	Products = append(Products, product)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(ID int) {
	for _, p := range Products {
		if p.ID == ID {
			fmt.Println(p)
		}
	}
}

var Products = []Product{}

func main() {
	Products = append(Products, Product{
		ID:          1,
		Name:        "Sabão",
		Price:       2.50,
		Description: "Lorem Ipsum",
		Category:    "Cheiroso",
	},
		Product{
			ID:          2,
			Name:        "Detergente",
			Price:       3.10,
			Description: "Lorem Ipsum",
			Category:    "Transparente",
		},
	)

	produtoTeste := Product{
		ID:          3,
		Name:        "Sabonete",
		Price:       5.60,
		Description: "Lorem Ipsum",
		Category:    "Do bom",
	}

	produtoTeste2 := Product{
		ID:          4,
		Name:        "Arroz",
		Price:       5.50,
		Description: "Lorem Ipsum",
		Category:    "São João",
	}

	produtoTeste.Save(produtoTeste)
	produtoTeste.GetAll()
	produtoTeste.Save(produtoTeste2)
	getById(4)
}
