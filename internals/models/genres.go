package models

import (
	"context"
	"time"
)

func (m *DBModel) GetGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "select id, genre_name, created_at, updated_at from genres"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre
	for rows.Next(){
		var genre Genre
		err := rows.Scan(&genre.ID, &genre.GenreName, &genre.CreatedAt, &genre.UpdatedAt)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}

	return genres, nil
}