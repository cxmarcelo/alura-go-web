package main

import (
	"net/http"
	"text/template"

	"github.com/cxmarcelo/alura-go-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}
