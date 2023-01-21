package models

import "time"

type Publicacao struct {
	ID         uint64    `json:"id,omitempty"`
	Titulo     string    `json:"titulo,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AutorID    uint64    `json:"autorId,omitempty"`
	AutorNick  uint64    `json:"autorNick,omitempty"`
	Curtidadas uint64    `json:"curtidadas"`
	CriadaEm   time.Time `json:"criadaEm,omitempty"`
}
