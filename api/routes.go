package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", s.getStatus)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.getMovie)
	r.HandlerFunc(http.MethodGet, "/v1/movies", s.getMovies)
	r.HandlerFunc(http.MethodGet, "/v1/genres", s.getGenres)
	r.HandlerFunc(http.MethodGet, "/v1/genres/:id/movies", s.getMoviesByGenre)
	r.HandlerFunc(http.MethodPost, "/v1/admin/movies", s.manageMovie)

	return s.enableCORS(r)
}