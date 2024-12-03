package routers

import (
	"golang-social-network-api/src/controllers"
	"net/http"
)

var routesPublications = []Routes{

	// Rota para criar uma publicação
	{
		URI:            	"/publications",
		Method:         	http.MethodPost,
		Function:       	controllers.CreatePublication,
		Authentication: 	true,
	},
	// Rota para buscar todas publicações do usuário
	{
		URI:            	"/publications",
		Method:         	http.MethodGet,
		Function:       	controllers.GetPublications,
		Authentication: 	true,
	},
	// Rota para buscar uma publicação
	{
		URI:            	"/publications/{id}",
		Method:         	http.MethodGet,
		Function:       	controllers.GetPublication,
		Authentication: 	true,
	},
	// Rota para atualizar publicação
	{
		URI:            	"/publications/{id}",
		Method:         	http.MethodPut,
		Function:       	controllers.UpdatePublication,
		Authentication: 	true,
	},
	// Rota para deletar publicação
	{
		URI:            	"/publications/{id}",
		Method:         	http.MethodDelete,
		Function:       	controllers.DeletePublication,
		Authentication: 	true,
	},
	// Rota para retornar todas as publicações de um usuário
	{
		URI:            	"/users/{id}/publications",
		Method:         	http.MethodGet,
		Function:       	controllers.GetPublicationByUser,
		Authentication: 	true,
	},
	// Rota para curtir publicações dos usuários
	{
		URI:            	"/publications/{id}/like",
		Method:         	http.MethodPost,
		Function:       	controllers.LikePublication,
		Authentication: 	true,
	},
	// Rota para descurtir publicações
	{
		URI:            	"/publications/{id}/dislike",
		Method:         	http.MethodPost,
		Function:       	controllers.DisLikePublication,
		Authentication: 	true,
	},
}