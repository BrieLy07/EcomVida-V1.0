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
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuarioID := params["id"]

		var direccion models.Direccion
		if err := json.NewDecoder(r.Body).Decode(&direccion); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
			return
		}

		direccion.UsuarioID = usuarioID

		if err := repo.Crear(direccion); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error al crear dirección"})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Dirección creada exitosamente"})
	}
}
