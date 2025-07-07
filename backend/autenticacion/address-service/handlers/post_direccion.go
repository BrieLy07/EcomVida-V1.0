package handlers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func PostDireccionHandler(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		usuarioID := params["id"]

		var direccion models.Direccion
		err := json.NewDecoder(r.Body).Decode(&direccion)
		if err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		direccion.UsuarioID = usuarioID

		err = repo.Crear(direccion)
		if err != nil {
			http.Error(w, "Error al crear dirección", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
