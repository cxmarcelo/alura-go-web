package main

import (
	"net/http"

	"github.com/cxmarcelo/alura-go-web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
