package database

func Migrar() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS roles (
			id SERIAL PRIMARY KEY,
			nombre VARCHAR(100) NOT NULL,
			descripcion TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS usuarios_roles (
			id SERIAL PRIMARY KEY,
			usuario_id BIGINT NOT NULL,
			rol_id BIGINT NOT NULL,
			FOREIGN KEY (rol_id) REFERENCES roles(id)
		);`,
	}

	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
