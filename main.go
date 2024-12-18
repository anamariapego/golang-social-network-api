package main

import (
	"fmt"
	"golang-social-network-api/src/config"
	"golang-social-network-api/src/router"
	"github.com/swaggo/http-swagger" 
	_ "golang-social-network-api/docs"
	"log"
	"net/http"
)

func main() {

	config.Load()

	fmt.Println("Executando a aplicação em: http://localhost:2468")

	// Criar um router
	r := router.Generate()

	// Rota para o Swagger UI
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler) 

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}