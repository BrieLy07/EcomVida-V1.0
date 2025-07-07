package controller

import (
	"net/http"
	"strconv"
	"warehouse-service/locator"
	"warehouse-service/model"

	"github.com/gin-gonic/gin"
)

func UpdateWarehouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var w model.Warehouse
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	repo := locator.Locator.GetWarehouseRepository()
	if err := repo.Update(id, w); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Almacén actualizado"})
}
