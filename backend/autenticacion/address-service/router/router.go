package router

import (
	"address-service/handlers"
	"address-service/repository"
	"database/sql"

	"github.com/gorilla/mux"
)

func ConfigurarRutas(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	repo := repository.NewDireccionRepository(db)

	r.HandleFunc("/users/{id}/addresses", handlers.ObtenerDirecciones(repo)).Methods("GET")
	r.HandleFunc("/users/{id}/addresses", handlers.CrearDireccion(repo)).Methods("POST")
	r.HandleFunc("/addresses/{id}", handlers.ActualizarDireccion(repo)).Methods("PUT")
	r.HandleFunc("/addresses/{id}", handlers.EliminarDireccion(repo)).Methods("DELETE")

	return r
}
