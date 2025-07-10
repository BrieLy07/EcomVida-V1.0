package controllers

import (
	"address-service/repository"
	"address-service/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ObtenerDireccion(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.ValidarToken(r)
	if err != nil {
		http.Error(w, `{"error":"Token inválido"}`, http.StatusUnauthorized)
		return
	}

	usuarioID := mux.Vars(r)["usuarioId"]
	tokenID := fmt.Sprintf("%v", (*claims)["id"])
	if tokenID != usuarioID {
		http.Error(w, `{"error":"Acceso denegado"}`, http.StatusForbidden)
		return
	}

	d, err := repository.ObtenerDireccion(usuarioID)
	if err != nil {
		http.Error(w, `{"error":"Dirección no encontrada"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}
