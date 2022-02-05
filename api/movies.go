package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/samirprakash/go-movies-server/internals/utils"
)

func (s *Server) getMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		s.logger.Println(errors.New("invalid movie id : "), err)
		utils.ErrorJSON(w, err)
		return
	}

	movie, err := s.models.DB.GetMovie(id)
	if err != nil {
		s.logger.Println("no movie found in the database")
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		s.logger.Println("error writing movie to response")
		utils.ErrorJSON(w, err)
		return
	}
}

func (s *Server) getMovies(w http.ResponseWriter, r *http.Request){
	movies, err := s.models.DB.GetMovies()
	if err != nil {
		s.logger.Println("error reading movies from the database")
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		s.logger.Println("error writing movies to response")
		utils.ErrorJSON(w, err)
		return
	}
}

func (s *Server) getMoviesByGenre(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		s.logger.Println(errors.New("invalid movie id : "), err)
		utils.ErrorJSON(w, err)
		return
	}

	movies, err := s.models.DB.GetMovies(id)
	if err != nil {
		s.logger.Println("error reading movies from the database")
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		s.logger.Println("error writing movies to response")
		utils.ErrorJSON(w, err)
		return
	}
}

func (s *Server) manageMovie(w http.ResponseWriter, r *http.Request){
	type jr struct {
		OK bool `json:"ok"`
	}

	ok := jr{
		OK: true,
	}

	err := utils.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}