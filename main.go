package main

import (
	"fmt"
	"golang-social-network-api/src/config"
	"golang-social-network-api/src/router"
	"log"
	"net/http"
)

func main() {

	config.Load()

	fmt.Println("Executando a aplicação em: http://localhost:2468")

	// Criar um router
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}