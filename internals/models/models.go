package models

import "time"

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

type Genre struct {
	ID        int       `json:"id,omitempty"`
	GenreName string    `json:"genre_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MovieGenre struct{
	ID        int       `json:"id,omitempty"`
	MovieID   int       `json:"movie_id,omitempty"`
	GenreID   int       `json:"genre_id,omitempty"`
	Genre     Genre     `json:"genre,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}