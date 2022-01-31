package main

import (
	"flag"
	"log"
	"os"

	"github.com/samirprakash/go-movies-server/api"
	"github.com/samirprakash/go-movies-server/internals/config"
)

func main(){
	var c config.Config

	flag.IntVar(&c.Port, "port", 4000, "port to start the server")
	flag.StringVar(&c.Env, "env", "development", "application environment (development|production")
	flag.Parse()

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	s := api.NewServer(c, l)
	err := s.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}