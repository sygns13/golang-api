package main

import (
	"log"
	"net/http"

	"github.com/sygns13/golang-api/authorization"
	"github.com/sygns13/golang-api/handler"
	"github.com/sygns13/golang-api/storage"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	//Aca cambiar por base de Datos
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Server Init, Listening on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
