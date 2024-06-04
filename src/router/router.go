package router

import (
	"github.com/gorilla/mux"
	"docker-volumes/src/router/routes"
)

func GenerateRouter() *mux.Router {

	r := mux.NewRouter()
	
	return routes.ConfigRoutes(r)
}
