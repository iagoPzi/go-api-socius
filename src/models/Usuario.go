package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *Usuario) Preparar(etapa string) error {
	if erro := user.validar(etapa); erro != nil {
		return erro
	}

	if erro := user.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (user *Usuario) validar(etapa string) error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório!")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório!")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório!")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido é inválido!")
	}
	if etapa == "cadastro" && user.Senha == "" {
		return errors.New("A senha é obrigatório!")
	}
	return nil
}

func (user *Usuario) formatar(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		senhaHash, erro := seguranca.Hash(user.Senha)
		if erro != nil {
			return erro
		}
		user.Senha = string(senhaHash)
	}
	return nil
}
