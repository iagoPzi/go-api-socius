package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}/unfollow",
		Metodo:     http.MethodPost,
		Funcao:     controllers.UnFollowUser,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguindo,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userid}/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
