package controllers

import (
	"cart-service/redis"
	"net/http"
)

func ClearCart(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Header.Get("X-User-Id")
	if usuarioID == "" {
		http.Error(w, "Falta el ID del usuario", http.StatusBadRequest)
		return
	}

	err := redis.Client.Del(redis.Ctx, usuarioID).Err()
	if err != nil {
		http.Error(w, "Error al eliminar el carrito", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"mensaje":"Carrito vaciado correctamente"}`))
}
