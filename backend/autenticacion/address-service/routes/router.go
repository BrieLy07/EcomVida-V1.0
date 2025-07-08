package routes

import (
	"address-service/handlers"
	"address-service/middleware"
	"address-service/repository"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigurarRutas(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	repo := repository.NewDireccionRepository(db)

	// Rutas protegidas por JWT
	api := r.PathPrefix("/").Subrouter()
	api.Use(middleware.VerificarJWT)

	direccionHandler := handlers.NewDireccionHandler(repo)
	api.HandleFunc("/users/{id}/addresses", direccionHandler.ObtenerDireccionesPorUsuario).Methods("GET")
	api.HandleFunc("/users/{id}/addresses", handlers.PostDireccionHandler(repo)).Methods("POST")
	api.HandleFunc("/addresses/{addressId}", handlers.PutDireccionHandler(repo)).Methods("PUT")
	api.HandleFunc("/addresses/{addressId}", handlers.DeleteDireccionHandler(repo)).Methods("DELETE")

	return r
}
