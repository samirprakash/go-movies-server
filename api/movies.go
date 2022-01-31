package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/samirprakash/go-movies-server/internals/models"
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

	movie := models.Movie{
		ID:          id,
		Title:       "sample movie",
		Description: "sample movie description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     120,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	utils.WriteJSON(w, http.StatusOK, movie, "movie")
}

func (s *Server) getAllMovies(w http.ResponseWriter, r *http.Request){

}