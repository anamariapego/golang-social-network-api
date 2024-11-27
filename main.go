package main

import (
	"fmt"
	"golang-social-network-api/src/config"
)

func main() {

	config.Load()

	fmt.Println("Executando a aplicação em: http://localhost:2468")

	// Criar um router
	//r := router.Generate()

	//log.Fatal(http.ListenAndServe(fmt.Sprint(":%d", config.Port), r))
}