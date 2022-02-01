package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", s.getStatus)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.getOneMovie)
	r.HandlerFunc(http.MethodGet, "/v1/movies", s.getAllMovies)

	return s.enableCORS(r)
}