package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}