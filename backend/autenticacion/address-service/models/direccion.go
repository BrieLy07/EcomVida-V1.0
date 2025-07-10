package models

type Direccion struct {
	ID           int    `json:"id,omitempty"`
	UsuarioID    string `json:"usuario_id"`
	Calle        string `json:"calle"`
	Ciudad       string `json:"ciudad"`
	Provincia    string `json:"provincia"`
	CodigoPostal string `json:"codigo_postal"`
}
