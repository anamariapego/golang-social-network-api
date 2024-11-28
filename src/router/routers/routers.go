package routers

import (
	"golang-social-network-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes representa a estrutura de rotas da API
type Routes struct {
	URI 			string
	Method 			string
	Function 		func(http.ResponseWriter, *http.Request)
	Authentication 	bool
}


// Config coloca todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {

	routes := routerUsers
	// routes = append(routes, routeLogin)
	// routes = append(routes, routesPublications...)

	for _, route := range routes {

		if route.Authentication {
			r.HandleFunc(route.URI, 
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r 
}