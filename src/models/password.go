package models

// Password representa o formato da requisição de atualização de senha do usuário
type Password struct {
	New 	string 	`json:"new"`
	Current string 	`json:"current"`
}