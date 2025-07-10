package handlers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func EliminarDireccion(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, `{"error": "ID inválido"}`, http.StatusBadRequest)
			return
		}

		err = repo.Eliminar(id)
		if err != nil {
			http.Error(w, `{"error": "Error al eliminar dirección"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"mensaje": "Dirección eliminada"})
	}
}
