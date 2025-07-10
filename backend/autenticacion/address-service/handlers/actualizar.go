package handlers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ActualizarDireccion(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, `{"error": "ID inválido"}`, http.StatusBadRequest)
			return
		}

		var direccion models.Direccion
		if err := json.NewDecoder(r.Body).Decode(&direccion); err != nil {
			http.Error(w, `{"error": "Datos inválidos"}`, http.StatusBadRequest)
			return
		}

		err = repo.Actualizar(id, direccion)
		if err != nil {
			http.Error(w, `{"error": "Error al actualizar dirección"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Dirección actualizada"})
	}
}
