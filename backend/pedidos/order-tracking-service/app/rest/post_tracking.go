package rest

import (
	"encoding/json"
	"net/http"
	"order-tracking-service/app/redis"

	"github.com/gorilla/mux"
)

type EstadoRequest struct {
	Estado string `json:"estado"`
}

func PostTrackingEstado(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var req EstadoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato inválido", http.StatusBadRequest)
		return
	}

	err := redis.GuardarEstado(orderID, req.Estado)
	if err != nil {
		http.Error(w, "Error guardando estado", http.StatusInternalServerError)
		return
	}

	redis.PublicarEstado(orderID, req.Estado)

	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Estado actualizado",
	})
}
