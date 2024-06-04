package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates(){
	templates = template.Must(template.ParseGlob("views/*.html"))
	// templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Renderiza uma página html na tela
func ExecTemplate(w http.ResponseWriter, template string, dados interface{}){
	templates.ExecuteTemplate(w, template, dados)
}