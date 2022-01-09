package handler

import (
	"net/http"

	"github.com/sygns13/golang-api/middleware"
)

// Route es una interface de rutas
func RoutePerson(mux *http.ServeMux, storage Storage) {

	h := newPerson(storage)

	// Registrando la ruta y el handler
	mux.HandleFunc("/api/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/api/v1/persons/update", h.update)
	mux.HandleFunc("/api/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/api/v1/persons/get-by-id", h.getByID)
	mux.HandleFunc("/api/v1/persons/get-all", middleware.Log(h.getAll))
	/* mux.HandleFunc("/person/", h.getByID)
	mux.HandleFunc("/person/all", h.getAll) */

}

//RouteLogin es una interface de rutas
func RouteLogin(mux *http.ServeMux, storage Storage) {

	h := newLogin(storage)

	// Registrando la ruta y el handler
	mux.HandleFunc("/api/v1/login", h.login)

}
