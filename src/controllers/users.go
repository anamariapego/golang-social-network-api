package controllers

import (
	"encoding/json"
	"fmt"
	"golang-social-network-api/src/auth"
	"golang-social-network-api/src/database"
	"golang-social-network-api/src/models"
	"golang-social-network-api/src/repositories"
	"golang-social-network-api/src/security"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser cria um novo usuário e o armazena no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro ao ler a requisição", http.StatusBadRequest)
		return
	}

	var user models.User
	if err = json.Unmarshal(corpusRequest, &user); err != nil {
		http.Error(w, "erro na conversão json para struct", http.StatusBadRequest)
		return
	}

	// Valida os valores 
	if err = user.Prepare("register"); err != nil {
		http.Error(w, fmt.Sprintf("erro na validação dos campos: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		log.Printf("erro na conexão com o banco de dados: %v\n", err)

		http.Error(w, "erro ao tentar conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para criar um novo usuário
	repository := repositories.NewReposUsers(db)
	userId, err := repository.Create(user)
	if err != nil {
		http.Error(w, "usuário com o mesmo email ou nick já existe", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso: id %d", userId)))
}

// GetUsers busca todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNikck := strings.ToLower(r.URL.Query().Get("user"))

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para buscar usuários
	repos := repositories.NewReposUsers(db)
	users, err := repos.Search(nameOrNikck)
	if err != nil {
		http.Error(w, "erro ao buscar o usuário no banco de dados", http.StatusBadRequest)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, users)

}

// GetUser busca uma usuário específico
func GetUser(w http.ResponseWriter, r *http.Request) {
	
	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para buscar um usuário
	repos := repositories.NewReposUsers(db)
	user, err := repos.SearchId(userId)
	if err != nil {
		http.Error(w, "erro ao buscar o id", http.StatusBadRequest)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, user)

}

// UpdateUser atualiza as informações de um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusInternalServerError)
		return
	}

	// Ler o usuário do token
	userIdToken, err := auth.ExtractUserId(r)
	fmt.Println(userIdToken)
	if err != nil {
		http.Error(w, "erro ao autenticar o usuário", http.StatusUnauthorized)
		return
	}

	if userId != userIdToken {
		http.Error(w, "acesso não autorizado", http.StatusForbidden)
		return
	}

	// Ler o corpo da requisição 
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "erro ao ler corpo da requisição", http.StatusInternalServerError)
		return
	}

	var user models.User
	if err = json.Unmarshal(corpusRequest, &user); err != nil {
		log.Printf("erro na conversão JSON: %v", err)
		http.Error(w, "erro na conversão json para struct", http.StatusBadRequest)
		return
	}

	if err = user.Prepare("edit"); err != nil {
		http.Error(w, "erro na edição das informações", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para atualizar os dados do usuário
	repos := repositories.NewReposUsers(db)
	if err = repos.Update(userId, user); err != nil {
		http.Error(w, "erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("informações atualizadas")

}

// DeleteUser deletar um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusInternalServerError)
		return
	}

	// Ler o usuário do token
	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, "erro ao autenticar o usuário", http.StatusUnauthorized)
		return
	}

	if userId != userIdToken {
		http.Error(w, "acesso não autorizado", http.StatusForbidden)
		return
	}

 	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para deletar um usuário
	repos := repositories.NewReposUsers(db)
	if err = repos.Delete(userId); err != nil {
		http.Error(w, "erro", http.StatusBadRequest)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("usuário deletado")
}

// FollowerUserd permite o usuário seguir outro usuário
func FollowerUserd(w http.ResponseWriter, r *http.Request) {
	
	// Extrair o id do usuário
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, "erro ao autenticar o usuário", http.StatusUnauthorized)
		return
	}

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusBadRequest)
		return
	}

	if followerId == userId {
		http.Error(w, "não é possível seguir você mesmo", http.StatusForbidden)
		return
	}
	
 	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para seguir um usuário
	repos := repositories.NewReposUsers(db)

	// Verifica se o usuário existe
	exists, err := repos.UserExists(userId)
	if err != nil || !exists {
		http.Error(w, "usuário a ser seguido não encontrado", http.StatusInternalServerError)
		return
	}

	if err = repos.StopFollower(userId, followerId); err != nil {
		http.Error(w, "erro ao seguir o usuário", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("usuário seguido com sucesso")
}

// StopFollowerUser permite parar de seguir um usuário
func StopFollowerUser(w http.ResponseWriter, r *http.Request) {
	
	// Extrair o id do usuário
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, "erro ao autenticar o usuário", http.StatusUnauthorized)
		return
	}

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusBadRequest)
		return
	}

	if followerId == userId {
		http.Error(w, "não é possível parar de seguir você mesmo", http.StatusForbidden)
		return
	}
    // Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para parar de seguir um usuário
	repos := repositories.NewReposUsers(db)

	// Verifica se o usuário existe
	exists, err := repos.UserExists(userId)
	if err != nil || !exists {
		http.Error(w, "usuário a ser deixado de seguir não encontrado", http.StatusNotFound)
		return
	}

	// Parar de seguir o usuário
	if err = repos.StopFollower(userId, followerId); err != nil {
		http.Error(w, "erro ao parar de seguir o usuário", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("usuário deixado de seguir com sucesso")

}

// GetFollowers busca todos os seguidores do usuário
func GetFollowers(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para buscar os seguidores do usuário
	repos := repositories.NewReposUsers(db)

	// Verifica se o usuário existe
	exists, err := repos.UserExists(userId)
	if err != nil || !exists {
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return
	}

	followers, err := repos.SearchFollowers(userId)
	if err != nil {
		http.Error(w, "erro ao buscar seguidores", http.StatusInternalServerError)
		return
	}

	//  Verifica se o usuário não tem seguidores
	if len(followers) == 0 {
		http.Error(w, "este usuário não possui seguidores", http.StatusOK)
		return
	}
	
	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, followers)
}

// GetFollowing busca todos os usuários que um usuário está seguindo
func GetFollowing(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para buscar os seguidores de um usuário
	repos := repositories.NewReposUsers(db)

	// Verifica se o usuário existe
	exists, err := repos.UserExists(userId)
	if err != nil || !exists {
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return
	}

	users, err := repos.SearchFollowing(userId)
	if err != nil {
		http.Error(w, "erro ao buscar usuários seguidos", http.StatusInternalServerError)
		return
	}

	// Se não houver usuários seguidos
	if len(users) == 0 {
		http.Error(w, "O usuário não segue nenhum usuário", http.StatusOK)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, users)
}

// UpdatePassword atualizar senha do usuário
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	// Extrai o id do usuário
	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, "erro ao autenticar o usuário", http.StatusUnauthorized)
		return
	}

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido", http.StatusBadRequest)
		return
	}

	if userId != userIdToken {
		http.Error(w, "não é possível atualizar senha de outro usuário", http.StatusForbidden)
		return
	}

	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	var password models.Password
	if err = json.Unmarshal(corpusRequest, &password); err != nil {
		http.Error(w, "erro no corpo da requisição", http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Repositório para atualizar a senha do usuário
	repos := repositories.NewReposUsers(db)

	// Verifica se o usuário existe
	exists, err := repos.UserExists(userId)
	if err != nil || !exists {
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return
	}

	passwordInDB, err := repos.SearchPassword(userId)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Verifica se as senhas são iguais
	if err = security.ValidatePassword(password.Current, passwordInDB); err != nil {
		http.Error(w, "A senha atual não condiz com a senha atual do banco", http.StatusUnauthorized)
		return
	}

	// Adicionar o hash na senha antes de alterar no banco
	passwordHash, err := security.FuncHash(password.New)
	if err != nil {
		http.Error(w, "erro ao adicionar o hash na senha", http.StatusBadRequest)
		return
	}
	
	// Atualização da senha no banco de dados
	if err = repos.UpdatePassword(userId, string(passwordHash)); err != nil {
		http.Error(w, "erro ao atualizar a senha", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("senha atualizada com sucesso")
}
