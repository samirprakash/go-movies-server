package main

import (
	"flag"
	"log"
	"os"

	"github.com/samirprakash/go-movies-server/api"
	"github.com/samirprakash/go-movies-server/internals/config"

	_ "github.com/lib/pq"
)

func main(){
	var c config.Config

	flag.IntVar(&c.Port, "port", 4000, "port to start the server")
	flag.StringVar(&c.Env, "env", "development", "application environment (development|production")
	flag.StringVar(&c.DB.DSN, "dsn", "postgres://root:secret@localhost:5432/go-movies?sslmode=disable", "postgres connection string")
	flag.Parse()

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	
	s := api.NewServer(c, l)

	db, err := s.OpenDB()
	if err != nil {
		l.Fatal("cannot connect to database")
	}
	defer db.Close()

	err = s.Start()
	if err != nil {
		l.Fatal("cannot start server", err)
	}
}