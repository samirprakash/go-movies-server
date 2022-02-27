package api

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type contextParamKey string

const (
	params contextParamKey = "params"
)

func (s *Server) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), params, ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (s *Server) routes() http.Handler {
	r := httprouter.New()
	secure := alice.New(s.checkToken)

	r.HandlerFunc(http.MethodPost, "/v1/signin", s.Signin)
	r.HandlerFunc(http.MethodGet, "/status", s.getStatus)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.getMovie)
	r.HandlerFunc(http.MethodDelete, "/v1/movies/:id", s.deleteMovie)
	r.HandlerFunc(http.MethodGet, "/v1/movies", s.getMovies)
	r.HandlerFunc(http.MethodGet, "/v1/genres", s.getGenres)
	r.HandlerFunc(http.MethodGet, "/v1/genres/:id/movies", s.getMoviesByGenre)

	r.POST("/v1/admin/movies", s.wrap(secure.ThenFunc(s.manageMovie)))

	return s.enableCORS(r)
}
