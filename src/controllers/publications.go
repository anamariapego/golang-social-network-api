package controllers

// import (
// 	"api/src/auth"
// 	"api/src/database"
// 	"api/src/models"
// 	"api/src/repositories"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// // Cria uma nova publicação
// func CreatePublication(w http.ResponseWriter, r *http.Request) {
// 	userId, err := auth.ExtractUserId(r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	corpusRequest, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
// 		return
// 	}

// 	var publication models.Publications
// 	if err = json.Unmarshal(corpusRequest, &publication); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	// Inserir o autor da publicação
// 	publication.AuthorId = userId

// 	// Validar valores
// 	if err = publication.Prepare(); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "Erro interno ao tentar conectar ao banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	repos := repositories.NewReposPublications(db)
// 	publication.Id, err = repos.Create(publication) 
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	// Codificar a publicação em JSON e escrever a resposta
// 	if err = json.NewEncoder(w).Encode(publication); err != nil {
// 		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
// 		return
// 	}
// }

// // GetPublications traz todas as publicações que aparecriam no feed do usuário
// func GetPublications(w http.ResponseWriter, r *http.Request) {
// 	userId, err := auth.ExtractUserId(r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	repos := repositories.NewReposPublications(db)
// 	publications, err := repos.Search(userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
	
// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	// Codificar a publicação em JSON e escrever a resposta
// 	if err = json.NewEncoder(w).Encode(publications); err != nil {
// 		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
// 		return
// 	}


// }

// // GetPublication traz uma única publicação
// func GetPublication(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)

// 	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	repos := repositories.NewReposPublications(db)
// 	publication, err := repos.SearchId(publicationId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Serializar a resposta em JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	err = json.NewEncoder(w).Encode(publication)
// 	if err != nil {
// 		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
// 		return
// 	}


// }

// // UpdatePublication atualiza os dados de uma publicação
// func UpdatePublication(w http.ResponseWriter, r *http.Request) {

// 	userId, err := auth.ExtractUserId(r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}
// 	params := mux.Vars(r)
// 	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	//
// 	repos := repositories.NewReposPublications(db)
// 	publicationInDB, err := repos.SearchId(publicationId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// verifica se os ids são iguais
// 	if publicationInDB.AuthorId != userId {
// 		http.Error(w, "Não é possível atualizar a publicação", http.StatusForbidden)
// 		return
// 	}

// 	corpusRequest, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
// 		return
// 	}

// 	var publication models.Publications
// 	if err = json.Unmarshal(corpusRequest, &publication); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err = publication.Prepare(); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err = repos.Update(publicationId, publication); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	// Codificar a publicação em JSON e escrever a resposta
// 	if err = json.NewEncoder(w).Encode(publication); err != nil {
// 		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
// 		return
// 	}

// }

// // DeletePublication deleta os dados de uma publicação
// func DeletePublication(w http.ResponseWriter, r *http.Request) {

// 	userId, err := auth.ExtractUserId(r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}
// 	params := mux.Vars(r)
// 	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	//
// 	repos := repositories.NewReposPublications(db)
// 	publicationInDB, err := repos.SearchId(publicationId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// verifica se os ids são iguais
// 	if publicationInDB.AuthorId != userId {
// 		http.Error(w, "Não é possível deletar a publicação", http.StatusForbidden)
// 		return
// 	}

// 	if err = repos.Delete(publicationId); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// }

// // GetPublicationByUser traz todas as publicacões de um usuário especifíco
// func GetPublicationByUser(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)
// 	userId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	repos := repositories.NewReposPublications(db)
// 	var publications []models.Publications
// 	publications, err = repos.SearchUser(userId) 
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	// Codificar a publicação em JSON e escrever a resposta
// 	if err = json.NewEncoder(w).Encode(publications); err != nil {
// 		http.Error(w, "Erro ao formatar a resposta em JSON", http.StatusInternalServerError)
// 		return
// 	}

// }

// // LikePublication é para curtir uma publicação
// func LikePublication(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)
// 	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	//
// 	repos := repositories.NewReposPublications(db)
// 	if err = repos.Like(publicationId); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNoContent)
// }

// // DisLikePublication é para descurtir uma publicação
// func DisLikePublication(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)
// 	publicationId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Conexão com o banco
// 	db, err := database.Connect()
// 	if err != nil {
// 		http.Error(w, "erro ao conectar com o banco de dados", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	//
// 	repos := repositories.NewReposPublications(db)
// 	if err = repos.DisLike(publicationId); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Configurar o cabeçalho como JSON e o status como 201 Created
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNoContent)
// }