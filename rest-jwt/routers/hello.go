package routers

import (
	"github.com/gorilla/mux"
	"github.com/kulcsartibor/rest-jwt/controllers"
	"github.com/kulcsartibor/rest-jwt/core/authentication"
	"github.com/urfave/negroni"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")
	return router
}
