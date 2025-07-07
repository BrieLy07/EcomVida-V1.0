package controller

import (
	"net/http"
	"strconv"
	"warehouse-service/locator"

	"github.com/gin-gonic/gin"
)

func DeleteWarehouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo := locator.Locator.GetWarehouseRepository()
	if err := repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Almacén eliminado"})
}
