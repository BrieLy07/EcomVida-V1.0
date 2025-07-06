package models

type Direccion struct {
	ID           int    `json:"id"`
	UsuarioID    string `json:"usuario_id"`
	Direccion    string `json:"direccion"`
	Ciudad       string `json:"ciudad"`
	Provincia    string `json:"provincia"`
	Pais         string `json:"pais"`
	CodigoPostal string `json:"codigo_postal"`
	Telefono     string `json:"telefono"`
}
