package rest

import (
	"encoding/json"
	"net/http"
	"order-tracking-service/app/redis"

	"github.com/gorilla/mux"
)

func GetTrackingEstado(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	estado, err := redis.ObtenerEstadoActual(orderID)
	if err != nil {
		http.Error(w, "No se encontró el estado del pedido", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id":     orderID,
		"estado": estado,
	})
}
