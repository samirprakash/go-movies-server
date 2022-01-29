package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "port to start the server")
	flag.StringVar(&cfg.env, "env", "development", "application environment (development|production")
	flag.Parse()

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request){
		s := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}

		j, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			log.Println("not able to marshal current statius", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Fatal("cannot start server")
	}
}