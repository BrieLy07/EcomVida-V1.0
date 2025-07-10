package controllers

import (
	"address-service/models"
	"address-service/repository"
	"address-service/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func ActualizarDireccion(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.ValidarToken(r)
	if err != nil {
		http.Error(w, `{"error":"Token inválido"}`, http.StatusUnauthorized)
		return
	}

	var d models.Direccion
	json.NewDecoder(r.Body).Decode(&d)

	usuarioID := fmt.Sprintf("%v", (*claims)["id"])
	if d.UsuarioID != usuarioID {
		http.Error(w, `{"error":"Acceso denegado"}`, http.StatusForbidden)
		return
	}

	err = repository.ActualizarDireccion(d)
	if err != nil {
		http.Error(w, `{"error":"No se pudo actualizar"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"mensaje":"Dirección actualizada"}`))
}
