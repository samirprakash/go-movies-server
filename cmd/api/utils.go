package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error{
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	j, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error){
	type je struct {
		Message string `json:"message"`
	}

	te := je{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, te, "error")
}