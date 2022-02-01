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

	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at from movies where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.Rating, &movie.MPAARating, &movie. CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *DBModel) GetAll() ([]*Movie, error){
	return nil, nil
}