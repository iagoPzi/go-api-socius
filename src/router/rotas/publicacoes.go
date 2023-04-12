package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoes,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{userId}/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoesPorUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CurtirPublicacao,
		RequerAuth: true,
	},

	{
		URI:        "/publicacoes/{publicacaoId}/descurtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.DescurtirPublicacao,
		RequerAuth: true,
	},
}
