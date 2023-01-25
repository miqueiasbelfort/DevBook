package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AutorID    uint64    `json:"autorId,omitempty"`
	AutorNick  string    `json:"autorNick,omitempty"`
	Curtidadas uint64    `json:"curtidadas"`
	CriadaEm   time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Prepara() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatorio e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O Conteudo é obrigatorio e não pode estar em branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
