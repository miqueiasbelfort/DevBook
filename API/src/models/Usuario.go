package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	usuario.formartar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome do usuário está em branco!")
	}
	if usuario.Nick == "" {
		return errors.New("O Nickname do usuário está em branco!")
	}
	if usuario.Email == "" {
		return errors.New("O Email do usuário está em branco!")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha do usuário está em branco!")
	}

	return nil
}

func (usuario *Usuario) formartar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
