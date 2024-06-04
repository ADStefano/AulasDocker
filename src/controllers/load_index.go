package controllers

import (
	"docker-volumes/src/utils"
	"net/http"
)

func LoadIndexTemplate(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "index.html", nil)
}
