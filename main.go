package main

import (
	"fmt"
	"log"
	"net/http"
	"docker-volumes/src/utils"
	"docker-volumes/src/router"
	"docker-volumes/src/config"
)

func main() {
	utils.LoadTemplates()
	r := router.GenerateRouter()

	fmt.Printf("PORTA: %d\nRodando APP...\n", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))
}
