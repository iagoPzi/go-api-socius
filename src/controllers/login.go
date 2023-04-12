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
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.Usuario
	if erro = json.Unmarshal(reqBody, &user); erro != nil {
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

	usuarioSalvoNoBanco, erro := repository.BuscarPorEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New("email e senha não condizem"))
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, user.Senha); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("email e senha não condizem"))
		return
	}
	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	userID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	responses.JSON(w, http.StatusOK, models.DadosAuth{ID: userID, Token: token})
}
