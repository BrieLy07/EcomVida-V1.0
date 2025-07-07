package controllers

import (
	"cart-service/memento"
	"cart-service/redis"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateRequest struct {
	Quantity int `json:"quantity"`
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Header.Get("X-User-Id")
	if usuarioID == "" {
		http.Error(w, "Falta el ID del usuario", http.StatusBadRequest)
		return
	}

	productID := mux.Vars(r)["productId"]
	if productID == "" {
		http.Error(w, "Falta el ID del producto", http.StatusBadRequest)
		return
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	val, err := redis.Client.Get(redis.Ctx, usuarioID).Result()
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	var carrito memento.CartState
	json.Unmarshal([]byte(val), &carrito)

	// Actualizar cantidad
	encontrado := false
	for i, item := range carrito.Items {
		if item.ProductID == productID {
			carrito.Items[i].Quantity = req.Quantity
			encontrado = true
			break
		}
	}

	if !encontrado {
		http.Error(w, "Producto no encontrado en el carrito", http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(carrito)
	redis.Client.Set(redis.Ctx, usuarioID, data, 0)

	w.Write([]byte(`{"mensaje":"Cantidad actualizada"}`))
}
