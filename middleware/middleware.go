package middleware

import (
	"log"
	"net/http"

	"github.com/sygns13/golang-api/authorization"
)

//Recibe una funcion de tipo handler y devuelve una funcion de tipo handler
func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Peticion %q, método %q", r.URL.Path, r.Method)
		f(w, r)
	}

}

//Authentication
func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}
		/*
			if token != "un-token-muy-seguro" {
				//responder con un error no autorizado
				//w.WriteHeader(http.StatusUnauthorized)
				forbidden(w, r)
				return
			} */

		f(w, r)
	}

}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}
