package handlers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CrearDireccion(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		usuarioID := mux.Vars(r)["id"]

		var direccion models.Direccion
		if err := json.NewDecoder(r.Body).Decode(&direccion); err != nil {
			http.Error(w, `{"error": "Datos inválidos"}`, http.StatusBadRequest)
			return
		}
		direccion.UsuarioID = usuarioID

		err := repo.Crear(direccion)
		if err != nil {
			http.Error(w, `{"error": "Error al crear dirección"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Dirección creada"})
	}
}
