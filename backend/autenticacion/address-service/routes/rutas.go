package routes

import (
	"address-service/controllers"

	"github.com/gorilla/mux"
)

func CargarRutas() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/address/me", controllers.ObtenerDireccion).Methods("GET")
	r.HandleFunc("/address", controllers.CrearDireccion).Methods("POST")
	r.HandleFunc("/address", controllers.ActualizarDireccion).Methods("PUT")
	r.HandleFunc("/address/me", controllers.EliminarDireccion).Methods("DELETE")

	return r
}
