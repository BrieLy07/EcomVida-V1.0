package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() *sql.DB {
	dsn := os.Getenv("DB_DSN") // ejemplo: usuario:password@tcp(host:puerto)/nombre_bd
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Base de datos no responde:", err)
	}

	return db
}
