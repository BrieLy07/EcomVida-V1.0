package handlers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DireccionHandler struct {
	Repo *repository.DireccionRepository
}

func NewDireccionHandler(repo *repository.DireccionRepository) *DireccionHandler {
	return &DireccionHandler{Repo: repo}
}

func (h *DireccionHandler) ObtenerDireccionesPorUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	usuarioID := vars["id"]

	direcciones, err := h.Repo.ObtenerPorUsuario(usuarioID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "No se pudieron obtener las direcciones"})
		return
	}

	json.NewEncoder(w).Encode(direcciones)
}
