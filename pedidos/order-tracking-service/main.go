package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	redisC *redis.Client
	upgr   = websocket.Upgrader{}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	redisC = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Servidor WebSocket en :8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgr.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer mensaje:", err)
			break
		}

		log.Printf("Mensaje recibido: %s", msg)

		err = redisC.Publish(ctx, "tracking", msg).Err()
		if err != nil {
			log.Println("Error al publicar en Redis:", err)
		}
	}
}
