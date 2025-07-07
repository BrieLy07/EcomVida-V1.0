package controllers

import (
	"cart-service/memento"
	"cart-service/redis"
	"encoding/json"
	"net/http"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Header.Get("X-User-Id")
	if usuarioID == "" {
		http.Error(w, "Falta el ID del usuario", http.StatusBadRequest)
		return
	}

	val, err := redis.Client.Get(redis.Ctx, usuarioID).Result()
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	var carrito memento.CartState
	json.Unmarshal([]byte(val), &carrito)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carrito)
}
