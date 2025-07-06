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
		params := mux.Vars(r)
		idStr := params["addressId"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var direccion models.Direccion
		err = json.NewDecoder(r.Body).Decode(&direccion)
		if err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		err = repo.Actualizar(id, direccion)
		if err != nil {
			http.Error(w, "Error al actualizar dirección", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
