package models

import (
	"database/sql"
	"time"
)

// Models defines that database models
type Models struct {
	DB DBModel
}

// NewModels returns Models with a database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie defines a movie
type Movie struct{
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Rating      int       `json:"rating"`
	MPAARating  string    `json:"mpaa_rating"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	MovieGenre map[int]string `json:"genres"`
}

// Genre defines the genres for movies
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MivieGenre defines the relationship between movie and genre
type MovieGenre struct{
	ID        int       `json:"-"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
}