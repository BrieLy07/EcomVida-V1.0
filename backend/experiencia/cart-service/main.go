package main

import (
	"cart-service/redis"
	"cart-service/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Cargar variables de entorno (opcional: usar librería si prefieres)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3027"
	}
	os.Setenv("REDIS_HOST", "localhost:6379") // puedes cambiarlo si apuntas a Redis en EC2

	// Inicializar Redis
	redis.InitRedis()

	// Configurar router
	r := mux.NewRouter()
	routes.RegisterCartRoutes(r)

	// Iniciar servidor
	log.Printf("Servidor iniciado en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
