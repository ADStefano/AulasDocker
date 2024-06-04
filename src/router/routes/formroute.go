package routes

import (
	"docker-volumes/src/controllers"
	"docker-volumes/src/models"
	"net/http"
)

var FormRoute = []models.Route{
	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: controllers.LoadIndexTemplate,
	},
	{
		URI:    "/save",
		Metodo: http.MethodPost,
		Funcao: controllers.SaveForm,
	},
	{
		URI:    "/load",
		Metodo: http.MethodGet,
		Funcao: controllers.LoadForm,
	},
}
