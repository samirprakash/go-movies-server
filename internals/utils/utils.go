package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error{
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

func ErrorJSON(w http.ResponseWriter, err error){
	type je struct {
		Message string `json:"message"`
	}

	te := je{
		Message: err.Error(),
	}

	WriteJSON(w, http.StatusBadRequest, te, "error")
}