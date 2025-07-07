package controllers

import (
	"cart-service/memento"
	"cart-service/redis"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RemoveCartItem(w http.ResponseWriter, r *http.Request) {
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

	val, err := redis.Client.Get(redis.Ctx, usuarioID).Result()
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	var carrito memento.CartState
	json.Unmarshal([]byte(val), &carrito)

	nuevaLista := []memento.CartItem{}
	for _, item := range carrito.Items {
		if item.ProductID != productID {
			nuevaLista = append(nuevaLista, item)
		}
	}
	carrito.Items = nuevaLista

	data, _ := json.Marshal(carrito)
	redis.Client.Set(redis.Ctx, usuarioID, data, 0)

	w.Write([]byte(`{"mensaje":"Producto eliminado del carrito"}`))
}
