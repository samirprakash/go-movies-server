package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/samirprakash/go-movies-server/internals/config"

	_ "github.com/lib/pq"
)

type Server struct {
	config config.Config
	logger *log.Logger
}

func NewServer(config config.Config, logger *log.Logger) *Server{
	s := &Server{
		config: config,
		logger: logger,
	}

	return s
}

func (s *Server) OpenDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", s.config.DB.DSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) Start() error{

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.config.Port),
		Handler:           s.routes(),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       time.Minute,
	}

	s.logger.Println("server starting on port : ", s.config.Port)

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}