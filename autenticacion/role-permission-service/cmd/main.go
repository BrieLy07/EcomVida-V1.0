package main

import (
	"log"
	"net/http"
	"os"

	"role-permission-service/internal/database"
	"role-permission-service/internal/handlers"
	"role-permission-service/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Conexión a base de datos
	if err := database.Conectar(); err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	// Ejecutar migraciones
	if err := database.Migrar(); err != nil {
		log.Fatal("Error ejecutando migración:", err)
	}

	// Configurar rutas
	r := mux.NewRouter()

	api := r.PathPrefix("/api/roles").Subrouter()
	api.Use(middleware.Autenticacion)

	api.HandleFunc("/", handlers.ObtenerRolesPorUsuario).Methods("GET")
	api.HandleFunc("/", handlers.CrearRol).Methods("POST")
	api.HandleFunc("/asignar", handlers.AsignarRol).Methods("POST")

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	log.Println("Servidor escuchando en puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
