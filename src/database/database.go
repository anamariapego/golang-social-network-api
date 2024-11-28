package database

import (
	"database/sql"
	"fmt"
	"golang-social-network-api/src/config"

	_ "github.com/lib/pq" // Driver de conexão com o Postgress
)

// Connect vai abrir a conexão com o banco de dados Postgress
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConnectDB)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}