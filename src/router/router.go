package router

import (
	"golang-social-network-api/src/router/routers"

	"github.com/gorilla/mux"
)

// Generate - retorna um router com as rotas configuradas
func Generate() *mux.Router {

	r := mux.NewRouter()
	return routers.Config(r)
}
