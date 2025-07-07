package repository

import (
	"database/sql"
	"warehouse-service/config"
	"warehouse-service/model"
)

type WarehouseRepository interface {
	FindAll() ([]model.Warehouse, error)
	FindByID(id int) (*model.Warehouse, error)
	Create(w model.Warehouse) error
	Update(id int, w model.Warehouse) error
	Delete(id int) error
}

type warehouseRepoImpl struct {
	db *sql.DB
}

func NewWarehouseRepository() WarehouseRepository {
	return &warehouseRepoImpl{
		db: config.DB,
	}
}

func (r *warehouseRepoImpl) FindAll() ([]model.Warehouse, error) {
	rows, err := r.db.Query("SELECT id, nombre, ubicacion, capacidad FROM almacenes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var almacenes []model.Warehouse
	for rows.Next() {
		var w model.Warehouse
		if err := rows.Scan(&w.ID, &w.Nombre, &w.Ubicacion, &w.Capacidad); err != nil {
			return nil, err
		}
		almacenes = append(almacenes, w)
	}
	return almacenes, nil
}

func (r *warehouseRepoImpl) FindByID(id int) (*model.Warehouse, error) {
	row := r.db.QueryRow("SELECT id, nombre, ubicacion, capacidad FROM almacenes WHERE id=$1", id)
	var w model.Warehouse
	if err := row.Scan(&w.ID, &w.Nombre, &w.Ubicacion, &w.Capacidad); err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *warehouseRepoImpl) Create(w model.Warehouse) error {
	_, err := r.db.Exec("INSERT INTO almacenes (nombre, ubicacion, capacidad) VALUES ($1, $2, $3)", w.Nombre, w.Ubicacion, w.Capacidad)
	return err
}

func (r *warehouseRepoImpl) Update(id int, w model.Warehouse) error {
	_, err := r.db.Exec("UPDATE almacenes SET nombre=$1, ubicacion=$2, capacidad=$3 WHERE id=$4", w.Nombre, w.Ubicacion, w.Capacidad, id)
	return err
}

func (r *warehouseRepoImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM almacenes WHERE id=$1", id)
	return err
}
