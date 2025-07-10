package controllers

import (
	"address-service/repository"
	"address-service/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func EliminarDireccion(w http.ResponseWriter, r *http.Request) {
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

	err = repository.EliminarDireccion(usuarioID)
	if err != nil {
		http.Error(w, `{"error":"No se pudo eliminar"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje":"Dirección eliminada"}`))
}
