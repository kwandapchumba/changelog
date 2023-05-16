package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	sqlc "github.com/kwandapchumba/prioritize/db/sqlc"
	"github.com/kwandapchumba/prioritize/response"
	"github.com/kwandapchumba/prioritize/token"
	"github.com/kwandapchumba/prioritize/utils"
)

var (
	internalError = "something went wrong"
)

// add new user
type addNewUserParams struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *BaseHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	var req addNewUserParams

	if err := d.Decode(&req); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			response.Error(w, internalError, http.StatusInternalServerError)
			return
		} else {
			log.Printf("error decoding request body to struct: %v", err)
			response.Error(w, internalError, http.StatusInternalServerError)
			return
		}
	}

	if req.Email == "" {
		log.Println("email is empty")
		response.Error(w, "email address cannot be empty", http.StatusBadRequest)
		return
	}

	if req.FullName == "" {
		log.Println("full name is empty")
		response.Error(w, "full name cannot be empty", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		log.Println("password is empty")
		response.Error(w, "password cannot be empty", http.StatusBadRequest)
		return
	}

	q := sqlc.New(h.db)

	emailExistsInDB, err := q.EmailExistsInDB(r.Context(), req.Email)
	if err != nil {
		log.Println(fmt.Errorf("could not check in email exists with error: %v", err))
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	if emailExistsInDB {
		log.Println(fmt.Errorf("email (%v) exists", req.Email))
		response.Error(w, "email is taken", http.StatusConflict)
		return
	}

	hashedPass, err := utils.HashPassword("")
	if err != nil {
		log.Println(err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	user, err := q.AddNewUser(r.Context(), sqlc.AddNewUserParams{
		UserID:       utils.RandomString(),
		Email:        req.Email,
		UserPassword: hashedPass,
	})
	if err != nil {
		log.Println(fmt.Errorf("AddNewUser: could not add new user with error: %v", err))
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	token, _, err := token.CreateToken(user.UserID, user.Email, time.Now().UTC(), config.AccessTokenDuration, false)
	if err != nil {
		log.Println("could not create access token")
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusCreated, token)
}
