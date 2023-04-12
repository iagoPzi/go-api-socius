package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.Usuario
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = user.Preparar("cadastro"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := db.ConnectDB()

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, erro := repository.Buscar(nomeOuNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, users)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.ConnectDB()

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, erro := repository.BuscarPorID(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	userIDToken, erro := auth.ExtrairUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusUnauthorized, errors.New("não é possível editar uma conta que não é sua"))
		return
	}

	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario models.Usuario
	if erro := json.Unmarshal(reqBody, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar("edicao"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if erro = repository.Atualizar(userID, usuario); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	userIDToken, erro := auth.ExtrairUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuario que não é seu"))
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	if erro = repository.Deletar(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {

	seguidorID, erro := auth.ExtrairUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	if erro = repository.Seguir(userID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, erro)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := auth.ExtrairUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possível não seguir você mesmo"))
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	if erro = repository.UnFollow(userID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, erro)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	seguidores, erro := repository.BuscarSeguidores(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, seguidores)

}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	users, erro := repository.BuscarSeguindo(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, users)

}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := auth.ExtrairUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["userid"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if userIDToken != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("não é possivel atualizar uma conta que não seja sua"))
		return
	}
	reqBody, erro := ioutil.ReadAll(r.Body)

	var senha models.Senha
	if err := json.Unmarshal(reqBody, &senha); err != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.ConnectDB()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repository := repositories.NewUsersRepository(db)
	senhaDB, err := repository.BuscarSenha(userID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if err := seguranca.VerificarSenha(senhaDB, senha.Atual); err != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("senha atual diferente da database"))
		return
	}

	senhaHash, err := seguranca.Hash(senha.Nova)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if err = repository.AtualizarSenha(userID, string(senhaHash)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
