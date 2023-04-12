package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtrairUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao
	if err := json.Unmarshal(reqBody, &publicacao); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	publicacao.AutorID = userID

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)
	publicacao.ID, err = repository.Criar(publicacao)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publicacao)
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtrairUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)
	publicacoes, err := repository.Buscar(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, publicacoes)

}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)
	publicacao, err := repository.BuscarPorID(publicacaoID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publicacao)
}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtrairUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewPublicacoesRepository(db)
	publicacaoDB, err := repository.BuscarPorID(publicacaoID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoDB.AutorID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel atualizar uma publicação que não é sua"))
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao models.Publicacao
	if err := json.Unmarshal(reqBody, &publicacao); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = publicacao.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Atualizar(publicacaoID, publicacao); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtrairUserID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewPublicacoesRepository(db)
	publicacaoDB, err := repository.BuscarPorID(publicacaoID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoDB.AutorID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel deletar uma publicação que não é sua"))
		return
	}

	if err := repository.Deletar(publicacaoID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)
	publicacoes, err := repository.BuscarPorUsuario(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publicacoes)

}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)

	if err := repository.Curtir(publicacaoID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoID, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicacoesRepository(db)

	if err := repository.Descurtir(publicacaoID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
