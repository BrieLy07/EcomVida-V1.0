package models

type Rol struct {
	ID          int64  `db:"id" json:"id"`
	Nombre      string `db:"nombre" json:"nombre"`
	Descripcion string `db:"descripcion" json:"descripcion"`
}

type Permiso struct {
	ID     int64  `db:"id" json:"id"`
	Nombre string `db:"nombre" json:"nombre"`
}

type UsuarioRol struct {
	ID        int64 `db:"id" json:"id"`
	UsuarioID int64 `db:"usuario_id" json:"usuario_id"`
	RolID     int64 `db:"rol_id" json:"rol_id"`
}
