package routes

import (
	"cart-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterCartRoutes(r *mux.Router) {
	r.HandleFunc("/cart", controllers.GetCart).Methods("GET")
	r.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	r.HandleFunc("/cart/{productId}", controllers.UpdateCartItem).Methods("PUT")
	r.HandleFunc("/cart/{productId}", controllers.RemoveCartItem).Methods("DELETE")
	r.HandleFunc("/cart/clear", controllers.ClearCart).Methods("DELETE")
}
