package models

import (
	"errors"
	"golang-social-network-api/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa a estrutura dos usuários
type User struct {
	Id  			uint64    `json:"id,omitempty"`
	NameUser 		string    `json:"nameUser,omitempty"`
	Nick  			string    `json:"nick,omitempty"`
	Email 			string    `json:"email,omitempty"`
	PasswordUser  	string 	  `json:"passwordUser,omitempty"`
	CreatedAt 	    time.Time `json:"createdAt,omitempty"`
}

// Métodos para validar validar parâmetros
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

// Método para verificar se todos os campos obrigatórios estão preenchidos
func (user *User) validate(stage string) error {
	// stage vai indicar se está numa etapa de cadastro ou edição de usuário
	if user.NameUser == "" {
		return errors.New("o nome é obrigatório")
	}
	if user.Nick == "" {
		return errors.New("o nick é obrigatório")
	}
	if user.Email == "" {
		return errors.New("o email é obrigatório")
	}
	// Valida se o formato do email é válido
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o email inserido é invalido")
	}

	// Se for cadastro e a senha estiver em branco a senha será obrigatória 
	if stage == "register" && user.PasswordUser == "" {
		return errors.New("a senha é obrigatório")
	}

	return nil
}

// Método para remover espaço em branco dos campos
func (user *User) format(stage string) error {
	user.NameUser = strings.TrimSpace(user.NameUser)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	// Se for cadastro de usuário realiza o hash da senha fornecida
	if stage == "register" {
		passwordHash, err := security.FuncHash(user.PasswordUser)
		if err != nil {
			return err
		}
		user.PasswordUser = string(passwordHash)
	}
	return nil
}