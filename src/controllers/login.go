package controllers

import (
	"encoding/json"
	"golang-social-network-api/src/auth"
	"golang-social-network-api/src/database"
	"golang-social-network-api/src/models"
	"golang-social-network-api/src/repositories"
	"golang-social-network-api/src/security"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login autentifica um usuário na api e retorna o token de acesso
func Login(w http.ResponseWriter, r *http.Request) {
	
	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro interno no servidor ao ler a requisição", http.StatusInternalServerError)
		return
	}

	// Armazena os dados da requisição numa variável
	var user models.User
	if err = json.Unmarshal(corpusRequest, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusBadRequest)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com os dados do usuário no banco de dados
	repos := repositories.NewReposUsers(db)
	userSalveDB, err := repos.SearchEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Valida a senha do usuário
	if err = security.ValidatePassword(user.PasswordUser, userSalveDB.PasswordUser); err != nil {
		http.Error(w, "email ou senha inválido", http.StatusUnauthorized)
		return
	}

	// Gera um token para o usuário 
	token, err := auth.CreateToken(userSalveDB.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	userId := strconv.FormatUint(userSalveDB.Id, 10)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.DataAuth{
		Id: userId,
		Token:  token,
	})
}