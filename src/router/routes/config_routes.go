package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigRoutes(r *mux.Router) *mux.Router {

	routes := FormRoute

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
