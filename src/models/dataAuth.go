package models

// DataAuth contém o token e o id do usuário autenticado
type DataAuth struct {
	Id		string	`json:"id"`
	Token 	string 	`json:"token"`
}