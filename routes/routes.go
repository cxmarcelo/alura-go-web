package routes

import (
	"net/http"

	"github.com/cxmarcelo/alura-go-web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
