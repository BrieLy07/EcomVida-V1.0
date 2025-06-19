package handlers

import (
	"net/http"

	"role-permission-service/internal/database"
	"role-permission-service/internal/middleware"
	"role-permission-service/internal/models"
	"role-permission-service/internal/utils"
)

func CrearRol(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol
	if err := utils.ParseJSON(r, &rol); err != nil {
		utils.Error(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	_, err := database.DB.Exec(`INSERT INTO roles (nombre, descripcion) VALUES ($1, $2)`, rol.Nombre, rol.Descripcion)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error al crear rol")
		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{"mensaje": "Rol creado"})
}

func AsignarRol(w http.ResponseWriter, r *http.Request) {
	var entrada models.UsuarioRol
	if err := utils.ParseJSON(r, &entrada); err != nil {
		utils.Error(w, http.StatusBadRequest, "Datos inválidos")
		return
	}

	_, err := database.DB.Exec(`INSERT INTO usuarios_roles (usuario_id, rol_id) VALUES ($1, $2)`, entrada.UsuarioID, entrada.RolID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error al asignar rol")
		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{"mensaje": "Rol asignado"})
}

func ObtenerRolesPorUsuario(w http.ResponseWriter, r *http.Request) {
	id := middleware.ObtenerUsuarioID(r)

	var roles []models.Rol
	err := database.DB.Select(&roles, `
		SELECT r.id, r.nombre, r.descripcion
		FROM roles r
		JOIN usuarios_roles ur ON ur.rol_id = r.id
		WHERE ur.usuario_id = $1
	`, id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error al obtener roles")
		return
	}

	utils.JSON(w, http.StatusOK, roles)
}
