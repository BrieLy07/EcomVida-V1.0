package handlers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetDireccionesHandler(repo *repository.DireccionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		usuarioID := params["id"]

		direcciones, err := repo.ObtenerPorUsuario(usuarioID)
		if err != nil {
			http.Error(w, "Error al obtener direcciones", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(direcciones)
	}
}
