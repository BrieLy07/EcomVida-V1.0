package handlers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PutDireccionHandler(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		idStr := params["addressId"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
			return
		}

		var direccion models.Direccion
		if err := json.NewDecoder(r.Body).Decode(&direccion); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
			return
		}

		if err := repo.Actualizar(id, direccion); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error al actualizar dirección"})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Dirección actualizada exitosamente"})
	}
}
