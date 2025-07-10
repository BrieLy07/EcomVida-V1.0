package repository

import (
	"address-service/database"
	"address-service/models"
)

func ObtenerDireccion(usuarioID string) (models.Direccion, error) {
	var d models.Direccion
	err := database.DB.QueryRow("SELECT id, usuario_id, calle, ciudad, provincia, codigo_postal FROM direcciones WHERE usuario_id = ?", usuarioID).
		Scan(&d.ID, &d.UsuarioID, &d.Calle, &d.Ciudad, &d.Provincia, &d.CodigoPostal)
	return d, err
}

func CrearDireccion(d models.Direccion) error {
	_, err := database.DB.Exec("INSERT INTO direcciones (usuario_id, calle, ciudad, provincia, codigo_postal) VALUES (?, ?, ?, ?, ?)",
		d.UsuarioID, d.Calle, d.Ciudad, d.Provincia, d.CodigoPostal)
	return err
}

func ActualizarDireccion(d models.Direccion) error {
	_, err := database.DB.Exec("UPDATE direcciones SET calle=?, ciudad=?, provincia=?, codigo_postal=? WHERE usuario_id=?",
		d.Calle, d.Ciudad, d.Provincia, d.CodigoPostal, d.UsuarioID)
	return err
}

func EliminarDireccion(usuarioID string) error {
	_, err := database.DB.Exec("DELETE FROM direcciones WHERE usuario_id=?", usuarioID)
	return err
}
