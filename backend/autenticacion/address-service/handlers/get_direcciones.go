package handlers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDireccionesHandler(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		usuarioID := params["id"]

		direcciones, err := repo.ObtenerPorUsuario(usuarioID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Error al obtener direcciones"})
			return
		}

		json.NewEncoder(w).Encode(direcciones)
	}
}
