package midddlewares

import (
	"golang-social-network-api/src/auth"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate verifica se o usuário que está fazendo a requisição está autenticado
// HandlerFunc é o mesmo que: func (w http.ResponseWriter, r *http.Request)
func Authenticate(next http.HandlerFunc) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValideteToken(r); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}