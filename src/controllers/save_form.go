package controllers

import (
	"docker-volumes/src/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Salva os dados do formulario em um arquivo txt dentro do docker
func SaveForm(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("nome")
	address := r.FormValue("endereco")
	age, err := strconv.Atoi(r.FormValue("idade"))

	if err != nil {
		http.Error(w, "Idade inválida", http.StatusBadRequest)
		return
	}

	log.Printf("Salvando os dados: Nome:%s, Endereço:%s, Idade:%d", name, address, age)

	if err := os.MkdirAll(config.BASE_DIR_PATH, 0777); err != nil {
		http.Error(w, "Não foi possível criar o diretório de armazenamento", http.StatusInternalServerError)
		return
	}

	files, _ := os.ReadDir(config.BASE_DIR_PATH)
	filesLen := len(files) + 1

	data := fmt.Sprintf("Nome:%s\nEndereço:%s\nIdade:%d\n", name, address, age)
	filePath := filepath.Join(config.BASE_DIR_PATH, fmt.Sprintf("%d.txt", filesLen))

	if err := os.WriteFile(filePath, []byte(data), 0644); err != nil {
		http.Error(w, "Não foi possível salvar os dados", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Dados salvos:\n%s", data)
}
