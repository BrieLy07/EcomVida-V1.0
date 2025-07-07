package controllers

import (
	"cart-service/memento"
	"cart-service/redis"
	"encoding/json"
	"net/http"
)

type AddRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.Header.Get("X-User-Id")
	if usuarioID == "" {
		http.Error(w, "Falta el ID del usuario", http.StatusBadRequest)
		return
	}

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Obtener carrito actual
	var carrito memento.CartState
	val, _ := redis.Client.Get(redis.Ctx, usuarioID).Result()
	if val != "" {
		json.Unmarshal([]byte(val), &carrito)
	} else {
		carrito = memento.CartState{UsuarioID: usuarioID, Items: []memento.CartItem{}}
	}

	// Verificar si el producto ya existe
	existe := false
	for i, item := range carrito.Items {
		if item.ProductID == req.ProductID {
			carrito.Items[i].Quantity += req.Quantity
			existe = true
			break
		}
	}
	if !existe {
		carrito.Items = append(carrito.Items, memento.CartItem{
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		})
	}

	// Guardar en Redis
	data, _ := json.Marshal(carrito)
	redis.Client.Set(redis.Ctx, usuarioID, data, 0)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"mensaje":"Producto agregado al carrito"}`))
}
