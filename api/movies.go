package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/samirprakash/go-movies-server/internals/utils"
)

func (s *Server) getOneMovie(w http.ResponseWriter, r *http.Request){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		s.logger.Println(errors.New("invalid movie id : "), err)
		utils.ErrorJSON(w, err)
		return
	}

	movie, err := s.models.DB.GetOne(id)
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

func (s *Server) getAllMovies(w http.ResponseWriter, r *http.Request){

}