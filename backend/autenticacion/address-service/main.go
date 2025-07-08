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
	log.Fatal(http.ListenAndServe(":"+port, habilitarCORS(r)))

}

// MIDDLEWARE CORS
func habilitarCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir solo el origen del frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://ecomvida-alb-973389764.us-east-1.elb.amazonaws.com")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
