package database

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

func Connect(connString string) *sql.DB {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Error abriendo conexión a DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error conectando a SQL Server: %v", err)
	}

	log.Println("Conexión a SQL Server exitosa")
	return db
}
