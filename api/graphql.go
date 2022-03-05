package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/samirprakash/go-movies-server/internals/models"
	"github.com/samirprakash/go-movies-server/internals/utils"

	"github.com/graphql-go/graphql"
)

var (
	movies []*models.Movie
	err    error
)

// gQL schema definition
var fields = graphql.Fields{
	"movie": &graphql.Field{
		Name:        "Get Movie",
		Description: "Get movie by ID",
		Type:        movieType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, movie := range movies {
					if movie.ID == id {
						return movie, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Name:        "Get all movies",
		Description: "Get all movies",
		Type:        graphql.NewList(movieType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return movies, nil
		},
	},
	"search": &graphql.Field{
		Name:        "Search for a movie",
		Description: "Search for a movie",
		Type:        graphql.NewList(movieType),
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var filteredList []*models.Movie
			search, ok := p.Args["titleContains"].(string)
			if ok {
				for _, searchedMovie := range movies {
					if strings.Contains(searchedMovie.Title, search) {
						filteredList = append(filteredList, searchedMovie)
					}
				}
			}
			return filteredList, nil
		},
	},
}

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:       "Movies",
		Interfaces: nil,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"release_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"runtime": &graphql.Field{
				Type: graphql.Int,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"mpaa_rating": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"poster": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (s *Server) moviesGraphQL(w http.ResponseWriter, r *http.Request) {
	movies, err = s.models.DB.GetMovies()
	if err != nil {
		utils.ErrorJSON(w, errors.New("gql: not able to get movies from database"), http.StatusInternalServerError)
		return
	}

	q, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, errors.New("gql: error reading request body"))
		return
	}

	query := string(q)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		utils.ErrorJSON(w, errors.New("gql: not able to create schema"))
		log.Println(err)
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		utils.ErrorJSON(w, fmt.Errorf("failed: %+v", res.Errors), http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, res, "response")
	if err != nil {
		utils.ErrorJSON(w, errors.New("gql: not able to write response"), http.StatusInternalServerError)
		return
	}
}
