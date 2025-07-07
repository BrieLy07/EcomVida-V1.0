package model

type Warehouse struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
	Capacidad int    `json:"capacidad"`
}
