package controller

import (
	"net/http"
	"strconv"
	"warehouse-service/locator"

	"github.com/gin-gonic/gin"
)

func GetWarehouseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo := locator.Locator.GetWarehouseRepository()
	result, err := repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Almacén no encontrado"})
		return
	}
	c.JSON(http.StatusOK, result)
}
