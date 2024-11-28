package database

import (
	"database/sql"
	"fmt"
	"golang-social-network-api/src/config"
)

// Connect vai abrir a conexão com o banco de dados Postgress
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postegres", config.StringConnectDB)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}