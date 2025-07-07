package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"order-tracking-service/app/redis"
	"order-tracking-service/app/rest"
	"order-tracking-service/app/ws"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando valores por defecto")
	}

	redis.InitRedis()

	router := mux.NewRouter()

	// REST
	router.HandleFunc("/orders/{id}/tracking", rest.GetTrackingEstado).Methods("GET")
	router.HandleFunc("/orders/{id}/tracking", rest.PostTrackingEstado).Methods("POST")

	// WebSocket
	router.HandleFunc("/ws/orders/{id}/track", ws.WebSocketTracking)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3012"
	}

	log.Printf("Servicio de seguimiento corriendo en el puerto %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
