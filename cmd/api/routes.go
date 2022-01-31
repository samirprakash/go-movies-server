package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	return r
}