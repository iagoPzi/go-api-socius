package models

import (
	"errors"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (p *Publicacao) Preparar() error {
	if erro := p.validar(); erro != nil {
		return erro
	}
	return nil
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("O titulo é obrigatório!")
	}
	if p.Conteudo == "" {
		return errors.New("O conteudo é obrigatório!")
	}

	return nil
}
