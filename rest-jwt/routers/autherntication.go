package routers

import (
	"github.com/kulcsartibor/rest-jwt/controllers"
	"github.com/kulcsartibor/rest-jwt/core/authentication"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc(
		"/token-auth",
		controllers.Login
	).Methods("POST")

	router.Handle(
		"/refresh-token-auth",
		negroni.New(
			negroni.HandleFunc(controllers.RefreshToken)
		).Methods("GET")
	)

	router.Handle(
		"/logout",
		negroni.New(
			negroni.HandleFunc(
				authentication.RequireTokenAuthentication
			),
			negroni.HandlerFunc(controllers.Logout))
		).Methods("GET")
	)

	return router
}