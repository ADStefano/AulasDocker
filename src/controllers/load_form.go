package controllers

import (
	"bufio"
	"docker-volumes/src/config"
	"docker-volumes/src/models"
	"docker-volumes/src/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	// "github.com/gorilla/mux"
)

// Carrega os dados do formulario salvo no arquivo com base no id do arquivo
func LoadForm(w http.ResponseWriter, r *http.Request) {

	fileParams := r.FormValue("id")

	fileId, err := strconv.ParseUint(fileParams, 10, 64)
	if err != nil {
		errTitle := fmt.Sprintf("Erro ao converter Id: %v", fileParams)
		log.Print(errTitle)
		http.Error(w, errTitle, http.StatusBadRequest)
		return
	}

	log.Printf("Carregando dados para arquivo com id: %d", fileId)

	filePath := filepath.Join(config.BASE_DIR_PATH, fmt.Sprintf("%d.txt", fileId))
	log.Printf("Caminho do arquivo: %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		errTitle := fmt.Sprintf("Erro ao recuperar os dados do arquivo(%d): %e", fileId, err)
		log.Print(errTitle)
		http.Error(w, errTitle, http.StatusInternalServerError)
		return
	}

	defer file.Close()

	var form models.Form

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "Nome"):
			nameString := strings.TrimSpace(line[len("Nome:"):])
			form.Name = nameString

		case strings.HasPrefix(line, "Endereço"):
			addressString := strings.TrimSpace(line[len("Endereço:"):])
			form.Address = addressString

		case strings.HasPrefix(line, "Idade"):
			ageStr := strings.TrimSpace(line[len("Idade:"):])
			intAge, err := strconv.Atoi(ageStr)
			age := int8(intAge)
			if err != nil {
				http.Error(w, fmt.Sprintf("Erro ao converter a idade: %e", err), http.StatusInternalServerError)
				return
			}
			form.Age = age
		}
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Erro ao ler o arquivo", http.StatusInternalServerError)
		return
	}

	utils.ExecTemplate(w, "load.html", form)
}
