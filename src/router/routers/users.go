package routers

import (
	"golang-social-network-api/src/controllers"
	"net/http"
)

// routerUsers rota de usuários
var routerUsers = []Routes{

	// Rota que cria usuários
	{
		URI: 				"/users",
		Method: 			http.MethodPost,
		Function: 			controllers.CreateUser,
		Authentication: 	false,
	},
	// Rota para buscar usuários
	{
		URI: 				"/users",
		Method: 			http.MethodGet,
		Function: 			controllers.GetUsers,
		Authentication: 	true,
	},
		// Rota para buscar um usuário
	{
		URI: 				"/users/{id}",
		Method: 			http.MethodGet,
		Function: 			controllers.GetUser,
		Authentication: 	true,
	},
	// Rota para atualizar informações de um usuário
	{
		URI: 				"/users/{id}",
		Method: 			http.MethodPut,
		Function: 			controllers.UpdateUser,
		Authentication: 	true,
	},
	// Rota para deletar usuário
	{
		URI: 				"/users/{id}",
		Method: 			http.MethodDelete,
		Function: 			controllers.DeleteUser,
		Authentication: 	true,
	},
	// Rota para seguir usuário
	{
		URI: 				"/users/{id}/follower",
		Method: 			http.MethodPost,
		Function: 			controllers.FollowerUserd,
		Authentication: 	true,
	},
	// Rota para parar de seguir um usuário
	{
		URI: 				"/users/{id}/stop-follower",
		Method: 			http.MethodPost,
		Function: 			controllers.StopFollowerUserd,
		Authentication: 	true,
	},
	// Rota para retornar todos os usuários que está seguindo
	{
		URI: 				"/users/{id}/followers",
		Method: 			http.MethodGet,
		Function: 			controllers.GetFollowers,
		Authentication:		true,
	},
	// Rota para retornar os usuários que um usuário especifico segue
	{
		URI: 				"/users/{id}/following",
		Method: 			http.MethodGet,
		Function: 			controllers.GetFollowing,
		Authentication: 	true,
	},
	// Rota para atualizar a senha do usuário
	{
		URI: 				"/users/{id}/update-password",
		Method: 			http.MethodPost,
		Function: 			controllers.UpdatePassword,
		Authentication: 	true,
	},
}