package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// DBModel defines the database model
type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetMovie(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// select one movie from movies table based on the movie_id and return the movie
	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at from movies where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.Rating, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get genres, if any
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
	defer rows.Close()

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

func (m *DBModel) GetMovies(genre ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf(`select id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at from movies %s order by title`, where)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// get genres, if any
		genreQuery := `select 
							mg.id, mg.movie_id, mg.genre_id, g.genre_name 
						from 
							movies_genres mg 
							left join genres g 
							on (g.id = mg.genre_id)
						where
							mg.movie_id = $1
						`
		genreRows, err := m.DB.QueryContext(ctx, genreQuery, movie.ID)
		if err != nil {
			return nil, err
		}

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
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
		genreRows.Close()

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *DBModel) InsertMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into movies 
				(title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at) 
				values 
				($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.CreatedAt,
		movie.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
