package controllers

import (
	"net/http"
	"role-permission-service/database"
	"role-permission-service/models"

	"github.com/gin-gonic/gin"
)

func CrearPermiso(c *gin.Context) {
	var permiso models.Permiso
	if err := c.ShouldBindJSON(&permiso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow("INSERT INTO permisos (nombre) VALUES ($1) RETURNING id", permiso.Nombre).Scan(&permiso.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el permiso"})
		return
	}

	c.JSON(http.StatusOK, permiso)
}

func ObtenerPermisos(c *gin.Context) {
	filas, err := database.DB.Query("SELECT id, nombre FROM permisos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener permisos"})
		return
	}
	defer filas.Close()

	var permisos []models.Permiso
	for filas.Next() {
		var permiso models.Permiso
		filas.Scan(&permiso.ID, &permiso.Nombre)
		permisos = append(permisos, permiso)
	}

	c.JSON(http.StatusOK, permisos)
}
