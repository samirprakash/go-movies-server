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
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Year        int       `json:"year,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Runtime     int       `json:"runtime,omitempty"`
	Rating      int       `json:"rating,omitempty"`
	MPAARating  string    `json:"mpaa_rating,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	MovieGenre []MovieGenre `json:"-"`
}

// Genre defines the genres for movies
type Genre struct {
	ID        int       `json:"id,omitempty"`
	GenreName string    `json:"genre_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// MivieGenre defines the relationship between movie and genre
type MovieGenre struct{
	ID        int       `json:"id,omitempty"`
	MovieID   int       `json:"movie_id,omitempty"`
	GenreID   int       `json:"genre_id,omitempty"`
	Genre     Genre     `json:"genre,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}