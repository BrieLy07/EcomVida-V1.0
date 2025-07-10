package repository

import (
	"address-service/models"
	"database/sql"
)

type DireccionRepository struct {
	DB *sql.DB
}

func NewDireccionRepository(db *sql.DB) *DireccionRepository {
	return &DireccionRepository{DB: db}
}

func (r *DireccionRepository) ObtenerPorUsuario(usuarioID string) ([]models.Direccion, error) {
	query := `SELECT id, usuario_id, direccion, ciudad, provincia, pais, codigo_postal, telefono 
	          FROM direcciones WHERE usuario_id = ?`
	rows, err := r.DB.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var direcciones []models.Direccion
	for rows.Next() {
		var d models.Direccion
		err := rows.Scan(&d.ID, &d.UsuarioID, &d.Direccion, &d.Ciudad, &d.Provincia, &d.Pais, &d.CodigoPostal, &d.Telefono)
		if err != nil {
			return nil, err
		}
		direcciones = append(direcciones, d)
	}
	return direcciones, nil
}

func (r *DireccionRepository) Crear(d models.Direccion) error {
	query := `
		INSERT INTO direcciones (usuario_id, direccion, ciudad, provincia, pais, codigo_postal, telefono)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, d.UsuarioID, d.Direccion, d.Ciudad, d.Provincia, d.Pais, d.CodigoPostal, d.Telefono)
	return err
}

func (r *DireccionRepository) Actualizar(id int, d models.Direccion) error {
	query := `
		UPDATE direcciones SET direccion=?, ciudad=?, provincia=?, pais=?, codigo_postal=?, telefono=?
		WHERE id=?`
	_, err := r.DB.Exec(query, d.Direccion, d.Ciudad, d.Provincia, d.Pais, d.CodigoPostal, d.Telefono, id)
	return err
}

func (r *DireccionRepository) Eliminar(id int) error {
	query := `DELETE FROM direcciones WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
