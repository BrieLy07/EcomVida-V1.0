package routes

import (
	"role-permission-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/roles", controllers.CrearRol)
	r.GET("/roles", controllers.ObtenerRoles)

	r.POST("/permisos", controllers.CrearPermiso)
	r.GET("/permisos", controllers.ObtenerPermisos)

	return r
}
