package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pascaldekloe/jwt"
	"github.com/samirprakash/go-movies-server/internals/utils"
)

func (s *Server) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(w, r)
	})
}

func (s *Server) checkToken(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2 {
			utils.ErrorJSON(w, errors.New("invalid auth header"))
			return
		}

		if headerParts[0] != "Bearer" {
			utils.ErrorJSON(w, errors.New("unauthorized - No Bearer"))
			return
		}

		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(s.config.Jwt.Secret))
		if err != nil {
			utils.ErrorJSON(w, errors.New("cannot check with HMACCheck"), http.StatusForbidden)
			return
		}

		if !claims.Valid(time.Now()) {
			utils.ErrorJSON(w, errors.New("unauthorized - token expired"), http.StatusForbidden)
			return
		}

		if !claims.AcceptAudience("mydomain.com") {
			utils.ErrorJSON(w, errors.New("unauthorized - invalid audience"), http.StatusForbidden)
			return
		}

		if claims.Issuer != "mydomain.com" {
			utils.ErrorJSON(w, errors.New("unauthorized - invalid issuer"), http.StatusForbidden)
			return
		}

		_, err = strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			utils.ErrorJSON(w, errors.New("unauthoried - cannot parse user ID"), http.StatusForbidden)
			return
		}

		next(w, r, ps)
	}
}
