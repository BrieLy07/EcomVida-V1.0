package main

import (
	"address-service/db"
	"address-service/handlers"
	"address-service/repository"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, se usarán variables de entorno del sistema")
	}

	db := db.Conectar()
	repo := repository.NewDireccionRepository(db)

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}/addresses", handlers.ObtenerDirecciones(repo)).Methods("GET")
	r.HandleFunc("/users/{id}/addresses", handlers.CrearDireccion(repo)).Methods("POST")
	r.HandleFunc("/addresses/{id}", handlers.ActualizarDireccion(repo)).Methods("PUT")
	r.HandleFunc("/addresses/{id}", handlers.EliminarDireccion(repo)).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Println("Servidor escuchando en el puerto", port)
	http.ListenAndServe(":"+port, r)
}
