package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel defines the database model
type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetOne(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// select one movie from movies table based on the movie_id and return the movie
	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at from movies where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.Rating, &movie.MPAARating, &movie. CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// combine movie and genre to populate genres in the selected movie
	query = `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name 
			from 
				movies_genres mg 
				left join genres g 
				on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
			`
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	
	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) GetAll() ([]*Movie, error){
	return nil, nil
}