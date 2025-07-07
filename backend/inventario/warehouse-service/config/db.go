package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConectarDB() error {
	url := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	DB = db
	fmt.Println("✅ Conexión a PostgreSQL exitosa")
	return nil
}
