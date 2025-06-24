package main

import (
	"log"
	"os"

	"role-permission-service/database"
	"role-permission-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar .env")
	}

	database.Conectar()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3004"
	}

	r.Run(":" + port)
}
