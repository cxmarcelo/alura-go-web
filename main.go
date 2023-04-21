package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta", Description: "Camiseta Azul", Price: 39, Quantity: 5},
		{Name: "Tenis", Description: "Confortável", Price: 139, Quantity: 3},
		{Name: "Fone", Description: "Muito bom", Price: 59, Quantity: 9},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
