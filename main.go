package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/samirprakash/go-movies-server/api"
	"github.com/samirprakash/go-movies-server/internals/config"

	_ "github.com/lib/pq"
)

func main() {
	var c config.Config

	flag.IntVar(&c.Port, "port", 4000, "port to start the server")
	flag.StringVar(&c.Env, "env", "development", "application environment (development|production")
	flag.StringVar(&c.DB.DSN, "dsn", "postgres://root:secret@localhost:5432/go_movies?sslmode=disable", "postgres connection string")
	flag.Parse()

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := OpenDB(c)
	if err != nil {
		l.Fatal("cannot connect to database")
	}
	defer db.Close()

	s := api.NewServer(c, l, db)

	err = s.Start()
	if err != nil {
		l.Fatal("cannot start server", err)
	}
}

func OpenDB(c config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", c.DB.DSN)
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
