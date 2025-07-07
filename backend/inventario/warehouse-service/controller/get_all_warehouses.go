package controller

import (
	"net/http"
	"warehouse-service/locator"

	"github.com/gin-gonic/gin"
)

func GetAllWarehouses(c *gin.Context) {
	repo := locator.Locator.GetWarehouseRepository()
	result, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener almacenes"})
		return
	}
	c.JSON(http.StatusOK, result)
}
