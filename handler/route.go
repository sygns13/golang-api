package handler

import "net/http"

// Route es una interface de rutas
func RoutePerson(mux *http.ServeMux, storage Storage) {

	h := newPerson(storage)

	// Registrando la ruta y el handler
	mux.HandleFunc("/api/v1/persons/create", h.create)
	mux.HandleFunc("/api/v1/persons/update", h.update)
	mux.HandleFunc("/api/v1/persons/delete", h.delete)
	mux.HandleFunc("/api/v1/persons/get-by-id", h.getByID)
	mux.HandleFunc("/api/v1/persons/get-all", h.getAll)
	/* mux.HandleFunc("/person/", h.getByID)
	mux.HandleFunc("/person/all", h.getAll) */

}
