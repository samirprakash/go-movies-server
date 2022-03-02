package api

import (
	"net/http"

	"github.com/samirprakash/go-movies-server/internals/utils"
)

func (s *Server) getGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := s.models.DB.GetGenres()
	if err != nil {
		s.logger.Println("error reading genres from database")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		s.logger.Println("erro writing genres to response")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}
