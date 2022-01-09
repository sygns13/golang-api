package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sygns13/golang-api/authorization"
	"github.com/sygns13/golang-api/model"
)

type login struct {
	storage Storage
}

func newLogin(storage Storage) login {
	return login{storage: storage}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Invalid JSON Estructura no Válida", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "Login no válido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "Error al generar token", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "Login exitoso", dataToken)
	responseJSON(w, http.StatusOK, response)

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "12345"
}
