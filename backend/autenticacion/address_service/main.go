package main

import (
	"log"
	"net/http"
	"os"

	"address-service/config"
	"address-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	db := config.ConectarBaseDatos()
	r := routes.ConfigurarRutas(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	log.Println("Servidor escuchando en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
