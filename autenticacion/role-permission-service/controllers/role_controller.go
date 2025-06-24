package controllers

import (
	"net/http"
	"role-permission-service/database"
	"role-permission-service/models"

	"github.com/gin-gonic/gin"
)

func CrearRol(c *gin.Context) {
	var rol models.Rol
	if err := c.ShouldBindJSON(&rol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.QueryRow("INSERT INTO roles (nombre) VALUES ($1) RETURNING id", rol.Nombre).Scan(&rol.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el rol"})
		return
	}

	c.JSON(http.StatusOK, rol)
}

func ObtenerRoles(c *gin.Context) {
	filas, err := database.DB.Query("SELECT id, nombre FROM roles")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener roles"})
		return
	}
	defer filas.Close()

	var roles []models.Rol
	for filas.Next() {
		var rol models.Rol
		filas.Scan(&rol.ID, &rol.Nombre)
		roles = append(roles, rol)
	}

	c.JSON(http.StatusOK, roles)
}
