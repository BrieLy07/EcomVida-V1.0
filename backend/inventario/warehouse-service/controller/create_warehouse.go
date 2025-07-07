package controller

import (
	"net/http"
	"warehouse-service/locator"
	"warehouse-service/model"

	"github.com/gin-gonic/gin"
)

func CreateWarehouse(c *gin.Context) {
	var w model.Warehouse
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	repo := locator.Locator.GetWarehouseRepository()
	if err := repo.Create(w); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensaje": "Almacén creado"})
}
