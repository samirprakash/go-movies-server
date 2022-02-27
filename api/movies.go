package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/samirprakash/go-movies-server/internals/models"
	"github.com/samirprakash/go-movies-server/internals/utils"
)

type jr struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (s *Server) getMovie(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) getMovies(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) getMoviesByGenre(w http.ResponseWriter, r *http.Request) {
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

type manageMovieRequestParams struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	MPAARating  string `json:"mpaa_rating"`
}

func (s *Server) manageMovie(w http.ResponseWriter, r *http.Request) {
	var req manageMovieRequestParams

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var movie models.Movie

	if req.ID != "0" {
		id, err := strconv.Atoi(req.ID)
		if err != nil {
			log.Println(err)
			utils.ErrorJSON(w, err)
			return
		}
		m, err := s.models.DB.GetMovie(id)
		if err != nil {
			log.Println(err)
			utils.ErrorJSON(w, err)
			return
		}
		movie = *m
		movie.UpdatedAt = time.Now()
	}

	movie.ID, err = strconv.Atoi(req.ID)
	if err != nil {
		log.Println(err)
		utils.ErrorJSON(w, err)
		return
	}
	movie.Title = req.Title
	movie.Description = req.Description
	movie.ReleaseDate, err = time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		log.Println(err)
		utils.ErrorJSON(w, err)
		return
	}
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, err = strconv.Atoi(req.Runtime)
	if err != nil {
		log.Println(err)
		utils.ErrorJSON(w, err)
		return
	}
	movie.Rating, err = strconv.Atoi(req.Rating)
	if err != nil {
		log.Println(err)
		utils.ErrorJSON(w, err)
		return
	}
	movie.MPAARating = req.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	if movie.ID == 0 {
		err = s.models.DB.InsertMovie(movie)
		if err != nil {
			log.Println(err)
			utils.ErrorJSON(w, err)
			return
		}
	} else {
		err = s.models.DB.UpdateMovie(movie)
		if err != nil {
			log.Println(err)
			utils.ErrorJSON(w, err)
			return
		}
	}

	ok := jr{
		OK: true,
	}

	err = utils.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

func (s *Server) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		s.logger.Println(errors.New("invalid movie id : "), err)
		utils.ErrorJSON(w, err)
		return
	}

	err = s.models.DB.DeleteMovie(id)
	if err != nil {
		s.logger.Println(errors.New("error deleting movie: "), err)
		utils.ErrorJSON(w, err)
		return
	}

	ok := jr{OK: true}

	err = utils.WriteJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		s.logger.Println(errors.New("error sending response while deleting a movie: "), err)
		utils.ErrorJSON(w, err)
		return
	}
}
