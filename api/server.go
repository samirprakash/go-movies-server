package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/samirprakash/go-movies-server/internals/config"
	"github.com/samirprakash/go-movies-server/internals/models"

	_ "github.com/lib/pq"
)

type Server struct {
	config config.Config
	logger *log.Logger
	models models.Models
}

func NewServer(config config.Config, logger *log.Logger, db *sql.DB) *Server{
	s := &Server{
		config: config,
		logger: logger,
		models: models.NewModels(db),
	}

	return s
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