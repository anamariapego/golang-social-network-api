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

// CreateUser cria usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "falha ao ler a requisição", http.StatusBadRequest)
		return
	}

	var user models.User
	if err = json.Unmarshal(corpusRequest, &user); err != nil {
		http.Error(w, "erro na conversão json para struct", http.StatusBadRequest)
		return
	}

	// Valida valores
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

	// Repositório para criar um novo usuário no banco de dados
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

	repos := repositories.NewReposUsers(db)
	users, err := repos.Search(nameOrNikck)
	if err != nil {
		http.Error(w, "erro ao buscar o usuário no banco de dados", http.StatusBadRequest)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
		return
	}

}

// GetUser busca uma usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "erro ao converter params para int", http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repos := repositories.NewReposUsers(db)
	user, err := repos.SearchId(userId)
	if err != nil {
		http.Error(w, "erro ao buscar o id", http.StatusBadRequest)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "erro ao formatar a resposta em JSON", http.StatusInternalServerError)
		return
	}

}

// UpdateUser atualiza as informações de um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// parâmetros da requisição
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "erro ao converter params para int", http.StatusInternalServerError)
		return
	}

	// Ler usuario do token
	userIdToken, err := auth.ExtractUserId(r)
	fmt.Println(userIdToken)
	if err != nil {
		http.Error(w, "token ausente ou inválido", http.StatusUnauthorized)
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

	// Atualizar no repositório
	repos := repositories.NewReposUsers(db)
	if err = repos.Update(userId, user); err != nil {
		http.Error(w, "erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	// Resposta com uma string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("informações atualizadas")

}

// DeleteUser deletar um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// parâmetros da requisição
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "erro ao converter params para int", http.StatusInternalServerError)
		return
	}

	// Ler usuário do token
	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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

	repos := repositories.NewReposUsers(db)
	if err = repos.Delete(userId); err != nil {
		http.Error(w, "erro", http.StatusBadRequest)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("usuário deletado")

}

// FollowerUserd permite o usuário seguir outro usuário
func FollowerUserd(w http.ResponseWriter, r *http.Request) {
	
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Id do parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	repos := repositories.NewReposUsers(db)
	exists, err := repos.UserExists(followerId)
	if err != nil {
		http.Error(w, "Erro ao verificar existência do seguidor", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Seguidor não existe", http.StatusBadRequest)
		return
	}

	if err = repos.Follower(userId, followerId); err != nil {
		log.Printf("erro ao seguir o usuário: %v", err)
		http.Error(w, "erro ao seguir o usuário", http.StatusBadRequest)
		return
	}

	// Serializar a resposta em JSON
	w.WriteHeader(http.StatusOK)
}

// StopFollowerUserd permite parar de seguir um usuário
func StopFollowerUserd(w http.ResponseWriter, r *http.Request) {
	
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Id do parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	repos := repositories.NewReposUsers(db)
	if err = repos.StopFollower(userId, followerId); err != nil {
		http.Error(w, "erro", http.StatusBadRequest)
		return
	}

	// Serializar a resposta em JSON
	w.WriteHeader(http.StatusOK)

}

// GetFollowers busca todos os seguidores do usuário
func GetFollowers(w http.ResponseWriter, r *http.Request) {

	// Id do parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repos := repositories.NewReposUsers(db)
	followers, err := repos.SearchFollowers(userId)
	if err != nil {
		http.Error(w, "erro ao buscar seguidores", http.StatusInternalServerError)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")

	// Se não houver seguidores
	if len(followers) == 0 {
		http.Error(w, "este usuário não possui seguidores", http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(followers); err != nil {
		http.Error(w, "erro ao formatar a resposta em JSON", http.StatusInternalServerError)
		return
	}

}

// GetFollowing busca todos os usuários que um usuário especifico está seguindo
func GetFollowing(w http.ResponseWriter, r *http.Request) {

	// Id do parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repos := repositories.NewReposUsers(db)
	users, err := repos.SearchFollowing(userId)
	if err != nil {
		http.Error(w, "erro ao buscar usuários seguidos", http.StatusInternalServerError)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "erro ao formatar a resposta em JSON", http.StatusInternalServerError)
		return
	}

}

// UpdatePassword atualizar senha do usuário
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	// Id do token
	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Id do parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userId != userIdToken {
		http.Error(w, "não é possível atualizar senha de outro usuário", http.StatusForbidden)
		return
	}

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

	repos := repositories.NewReposUsers(db)
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

	if err = repos.UpdatePassword(userId, string(passwordHash)); err != nil {
		http.Error(w, "erro ao atualizar a senha", http.StatusInternalServerError)
		return
	}

	// Serializar a resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`"Senha atualizada com sucesso"`))

}

