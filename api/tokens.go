package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"github.com/samirprakash/go-movies-server/internals/models"
	"github.com/samirprakash/go-movies-server/internals/utils"
	"golang.org/x/crypto/bcrypt"
)

var user = models.User{
	ID:       1,
	Email:    "me@here.com",
	Password: string(generateEncryptedPassword("password")),
}

func generateEncryptedPassword(p string) []byte {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hashedPassword
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	hashedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprintf(user.Email)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(s.config.Jwt.Secret))
	if err != nil {
		utils.ErrorJSON(w, errors.New("error signing"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, string(jwtBytes), "response")
	if err != nil {
		utils.ErrorJSON(w, errors.New("error sending signin response"))
		return
	}
}
