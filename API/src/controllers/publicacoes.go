package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/responstas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responstas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responstas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioId

	if erro := publicacao.Prepara(); erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusCreated, publicacao)

}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responstas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusOK, publicacoes)
}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPublicacaoPorID(publicacaoID)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusOK, publicacao)

}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responstas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPublicacaoPorID(publicacaoID)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		responstas.Erro(w, http.StatusForbidden, errors.New("Somente o usu??rio que publicou a altera????o pode edita-la"))
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responstas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Prepara(); erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusNoContent, nil)

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responstas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPublicacaoPorID(publicacaoID)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		responstas.Erro(w, http.StatusForbidden, errors.New("Somente o usu??rio que publicou pode apagar a publica????o"))
		return
	}

	if erro = repositorio.Deletar(publicacaoID); erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusNoContent, nil)

}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usaurioId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Curtir(publicacaoID); erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusNoContent, nil)
}

func DesCurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametro["publicacaoId"], 10, 64)
	if erro != nil {
		responstas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	if erro = repositorio.Descurtir(publicacaoID); erro != nil {
		responstas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responstas.JSON(w, http.StatusNoContent, nil)
}
