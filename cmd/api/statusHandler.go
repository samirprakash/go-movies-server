package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) getStatus(w http.ResponseWriter, r *http.Request){
	s := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	j, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		app.logger.Println("not able to marshal current status", err)
	}

	app.writeJSON(w, http.StatusOK, j, "status")
}