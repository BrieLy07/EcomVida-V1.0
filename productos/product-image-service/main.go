package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"product-image-service/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, se usarán valores por defecto.")
	}

	r := mux.NewRouter()

	r.HandleFunc("/imagenes", handlers.SubirImagen).Methods("POST")
	r.HandleFunc("/imagenes/{nombre}", handlers.ObtenerImagen).Methods("GET")

	log.Println("Servidor corriendo en puerto 8080...")
	http.ListenAndServe(":8080", r)
}
