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

// Login é responsável por autentificar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err = json.Unmarshal(corpusRequest, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusBadRequest)
		return
	}
	defer db.Close()

	repos := repositories.NewReposUsers(db)
	userSalveDB, err := repos.SearchEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = security.ValidatePassword(user.PasswordUser, userSalveDB.PasswordUser); err != nil {
		http.Error(w, "email ou senha inválido", http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(userSalveDB.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userId := strconv.FormatUint(userSalveDB.Id, 10)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.DataAuth{
		Id: userId,
		Token:  token,
	})
}