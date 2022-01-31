package api

import (
	"net/http"

	"github.com/samirprakash/go-movies-server/internals/utils"
)

const version = "1.0.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (s *Server) getStatus(w http.ResponseWriter, r *http.Request){
	status := AppStatus{
		Status:      "Available",
		Environment: s.config.Env,
		Version:     version,
	}

	 utils.WriteJSON(w, http.StatusOK, status, "status")
}