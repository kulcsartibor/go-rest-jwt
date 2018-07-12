package main

import (
	"github.com/kulcsartibor/rest-jwt/routers"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRouters()
	n := negroni.Classic()
	http.ListenAndServe(":8000", n)
}