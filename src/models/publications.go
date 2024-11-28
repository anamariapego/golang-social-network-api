package models

import (
	"errors"
	"strings"
	"time"
)

// Publications representa a estrutura das publicações
type Publications struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Text       string    `json:"text,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Métodos para validar parâmetros
func (publication *Publications) Prepare() error {
	if err := publication.valide(); err != nil {
		return err
	}

	publication.format()
	return nil 
}

func (publication *Publications) valide() error {
	if publication.Title == "" {
		return errors.New("o título é obrigatório")
	}

	if publication.Text == "" {
		return errors.New("o texto é obrigatório")
	}

	return nil
}

// Método para remover espaço em branco dos campos
func (publication *Publications) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Text = strings.TrimSpace(publication.Text)
}