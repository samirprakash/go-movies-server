package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
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
		utils.ErrorJSON(w, err, http.StatusNotFound)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		s.logger.Println("error writing movie to response")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func (s *Server) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := s.models.DB.GetMovies()
	if err != nil {
		s.logger.Println("error reading movies from the database")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		s.logger.Println("error writing movies to response")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
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
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		s.logger.Println("error writing movies to response")
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
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
			utils.ErrorJSON(w, err, http.StatusNotFound)
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

	if movie.Poster == "" {
		movie = getPoster(movie)
	}

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

func (s *Server) deleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
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
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func getPoster(movie models.Movie) models.Movie {
	type tmdb struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			Popularity       float64 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date,omitempty"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	}

	client := &http.Client{}
	key := "a9c77f4cfde39e7a5f62c9c5dee25ee9"
	url := "https://api.themoviedb.org/3/search/movie/?api_key=" + key + "&query=" + url.QueryEscape(movie.Title)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return movie
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return movie
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return movie
	}

	var tmdbResponse tmdb
	json.Unmarshal(bodyBytes, &tmdbResponse)

	if len(tmdbResponse.Results) > 0 {
		movie.Poster = tmdbResponse.Results[0].PosterPath
	}

	return movie
}
