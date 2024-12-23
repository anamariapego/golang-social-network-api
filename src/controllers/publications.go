package controllers

import (
	"encoding/json"
	"golang-social-network-api/src/auth"
	"golang-social-network-api/src/database"
	"golang-social-network-api/src/models"
	"golang-social-network-api/src/repositories"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication cria uma nova publicação do usuário autenticado
func CreatePublication(w http.ResponseWriter, r *http.Request) {

	// Extrai o id do usuário do token
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Estrutura da publicação
	var publication models.Publications
	if err = json.Unmarshal(corpusRequest, &publication); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Define o id do autor da publicação
	publication.AuthorId = userId

	// Valida os campos
	if err = publication.Prepare(); err != nil {
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

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)
	publication.Id, err = repos.Create(publication) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, publication)
}

// GetPublications retorna todas as publicações que aparecem no feed do usuário
func GetPublications(w http.ResponseWriter, r *http.Request) {
	
	// Extrai o id do usuário do token
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)
	publications, err := repos.SearchPublications(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, publications)
}

// GetPublication retorna uma única publicação com base no Id da publicação
func GetPublication(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id da publicação inválido", http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)

	// Verifica se a publicação existe
	exists, err := repos.ExistPublications(publicationId)
	if err != nil || !exists {
		http.Error(w, "publicação não encontrada", http.StatusNotFound)
		return
	}

	publication, err := repos.SearchPublicationsId(publicationId)
	if err != nil {
		http.Error(w, "erro ao buscar a publicação", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, publication)
}

// UpdatePublication atualiza os dados de uma publicação
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	// Extrai o id do usuário do token
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id da publicação inválido", http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)

	// Verifica se a publicação existe
	exists, err := repos.ExistPublications(publicationId)
	if err != nil || !exists {
		http.Error(w, "publicação não encontrada", http.StatusNotFound)
		return
	}

	publicationInDB, err := repos.SearchPublicationsId(publicationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verifica se os ids são iguais
	if publicationInDB.AuthorId != userId {
		http.Error(w, "não é possível atualizar a publicação de outro usuário", http.StatusForbidden)
		return
	}

	// Lê o corpo da requisição
	corpusRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Estrutura da publicação
	var publication models.Publications
	if err = json.Unmarshal(corpusRequest, &publication); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = publication.Prepare(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = repos.Update(publicationId, publication); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	JsonResponse(w, http.StatusCreated, publication)
}

// DeletePublication deleta os dados de uma publicação
func DeletePublication(w http.ResponseWriter, r *http.Request) {

	// Extrai o id do usuário do token
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id da publicação inválido", http.StatusInternalServerError)
		return
	}

	// Conexão com o banco de dados
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)

	// Verifica se a publicação existe
	exists, err := repos.ExistPublications(publicationId)
	if err != nil || !exists {
		http.Error(w, "publicação não encontrada", http.StatusNotFound)
		return
	}

	publicationInDB, err := repos.SearchPublicationsId(publicationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verifica se os ids são iguais
	if publicationInDB.AuthorId != userId {
		http.Error(w, "você não tem permissão para deletar esta publicação", http.StatusForbidden)
		return
	}

	if err = repos.Delete(publicationId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("publicação deletada com sucesso")
}

// GetPublicationByUser retorna todas as publicações de um usuário específico
func GetPublicationByUser(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id do usuário inválido.", http.StatusBadRequest)
		return
	}

	// Conexão com o banco
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Estrutura de publicação
	var publications []models.Publications

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)
	publications, err = repos.SearchUser(userId) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho como JSON e o status como 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Codificar a publicação em JSON e escrever a resposta
	if err = json.NewEncoder(w).Encode(publications); err != nil {
		http.Error(w, "erro ao formatar a resposta em JSON", http.StatusInternalServerError)
		return
	}
}

// LikePublication é para curtir uma publicação
func LikePublication(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id da publicação inválido", http.StatusBadRequest)
		return
	}

	// Conexão com o banco
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)

	// Verifica se a publicação existe
	exists, err := repos.ExistPublications(publicationId)
	if err != nil || !exists {
		http.Error(w, "publicação não encontrada", http.StatusNotFound)
		return
	}
	if err = repos.Like(publicationId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho como JSON e o status como 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// DisLikePublication é para descurtir uma publicação
func DisLikePublication(w http.ResponseWriter, r *http.Request) {

	// Obtém o id dos parâmetros
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "id da publicação inválido", http.StatusBadRequest)
		return
	}

	// Conexão com o banco
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Cria um repositório para interagir com o banco de dados
	repos := repositories.NewReposPublications(db)

	// Verifica se a publicação existe
	exists, err := repos.ExistPublications(publicationId)
	if err != nil || !exists {
		http.Error(w, "publicação não encontrada", http.StatusNotFound)
		return
	}
	if err = repos.DisLike(publicationId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar o cabeçalho como JSON e o status como 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}


// JsonResponse envia uma resposta JSON padronizada ao cliente.
func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "erro ao formatar a resposta em JSON", http.StatusInternalServerError)
    }
}