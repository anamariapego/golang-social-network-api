package routers

import (
	"golang-social-network-api/src/controllers"
	"net/http"
)

var routeLogin = Routes{
	URI:            	"/login",
	Method:         	http.MethodPost,
	Function:       	controllers.Login,
	Authentication: 	false,
}
