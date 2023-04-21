package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/cxmarcelo/alura-go-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.InsertProduct(name, description, convertedPrice, convertedQuantity)

	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	temp.ExecuteTemplate(w, "New", nil)
}
