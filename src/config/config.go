package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o banco de dados Postegress
	StringConnectDB string = ""

	// Porta que API está rodando
	Port int = 0

	// Chave para assinar o token
	SecretKey []byte
)

// Load vai inicializar as variáveis de ambiente
func Load() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Definindo uma porta padrão se enão for encontrada no arquivo .env
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	// Converte a string para inteiro
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Printf("erro ao converter DB_PORT para inteiro: %v\n", err)
		return
	}

	// Cria a string de conexão para o banco de dados
	StringConnectDB = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), dbPort, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
