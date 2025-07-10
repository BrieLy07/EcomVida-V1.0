package handlers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ObtenerDirecciones(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		usuarioID := mux.Vars(r)["id"]

		direcciones, err := repo.ObtenerPorUsuario(usuarioID)
		if err != nil {
			http.Error(w, `{"error": "Error al obtener direcciones"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(direcciones)
	}
}
