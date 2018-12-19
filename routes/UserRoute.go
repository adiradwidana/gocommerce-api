package routes

import (
	"github.com/gorilla/mux"
	"hello/controllers"
	"hello/repositories"
	"net/http"
)
// Collect all routes of users
type UserRoutes []BaseRoute

var UserController = &controllers.UserController{UserRepository: repositories.UserRepository{}}
var routes = UserRoutes{
	BaseRoute{
		"Index",
		"GET",
		"/users",
		UserController.Index,
	},
}

func UserRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	return router;
}
