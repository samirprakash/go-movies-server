package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) wrap(next http.Handler) httprouter.Handle {
	return s.checkToken(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next.ServeHTTP(w, r)
	})
}

func (s *Server) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/v1/graphql", s.moviesGraphQL)
	r.HandlerFunc(http.MethodPost, "/v1/signin", s.signin)
	r.HandlerFunc(http.MethodGet, "/status", s.getStatus)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", s.getMovie)
	r.HandlerFunc(http.MethodGet, "/v1/movies", s.getMovies)
	r.HandlerFunc(http.MethodGet, "/v1/genres", s.getGenres)
	r.HandlerFunc(http.MethodGet, "/v1/genres/:id/movies", s.getMoviesByGenre)

	r.POST("/v1/admin/movies", s.wrap(http.HandlerFunc(s.manageMovie)))
	r.DELETE("/v1/movies/:id", s.checkToken(s.deleteMovie))

	return s.enableCORS(r)
}
