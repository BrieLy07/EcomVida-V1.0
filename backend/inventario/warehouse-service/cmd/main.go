package main

import (
	"log"
	"os"
	"warehouse-service/config"
	"warehouse-service/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar .env, se usará el entorno del sistema.")
	}

	// Conectar a PostgreSQL
	if err := config.ConectarDB(); err != nil {
		log.Fatal("Error al conectar a PostgreSQL:", err)
	}

	router := gin.Default()

	// Endpoints
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"mensaje": "Warehouse-service activo"})
	})

	router.GET("/warehouses", controller.GetAllWarehouses)
	router.GET("/warehouses/:id", controller.GetWarehouseByID)
	router.POST("/warehouses", controller.CreateWarehouse)
	router.PUT("/warehouses/:id", controller.UpdateWarehouse)
	router.DELETE("/warehouses/:id", controller.DeleteWarehouse)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3017"
	}
	router.Run(":" + port)
}
