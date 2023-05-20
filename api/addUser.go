package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	sqlc "github.com/kwandapchumba/changelog/db/sqlc"
	"github.com/kwandapchumba/changelog/response"
	"github.com/kwandapchumba/changelog/token"
	"github.com/kwandapchumba/changelog/utils"
)

var (
	internalError = "something went wrong"
)

type addNewUserRequestParams struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s addNewUserRequestParams) validate() error {
	return validation.ValidateStruct(&s, validation.Field(&s.FullName, validation.Required.Error("full name required")), validation.Field(&s.Password, validation.Required.Error("password is required"), validation.Length(6, 0).Error("password must be at least six characters long")), validation.Field(&s.Email, is.Email.Error("email must be valid"), validation.Length(3, 0).Error("email must be at least three characters long")))
}

func (h *BaseHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	var req addNewUserRequestParams

	if err := d.Decode(&req); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			response.Error(w, internalError, http.StatusInternalServerError)
			return
		} else {
			log.Printf("error decoding request body to struct %v", err)
			response.Error(w, internalError, http.StatusInternalServerError)
			return
		}
	}

	if err := req.validate(); err != nil {
		log.Printf("adduser: %v", err)
		response.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	q := sqlc.New(h.db)

	otp, err := q.GetOtpByEmail(r.Context(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("adduser, getoptbymail: otp not found")
			response.Error(w, "otp not found", http.StatusForbidden)
			return
		}

		log.Printf("adduser: getotpbyemail: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	if !otp.Verified.Valid {
		log.Printf("adduser: email has not been verified")
		response.Error(w, "email has not been verified", http.StatusForbidden)
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
		FullName:     req.FullName,
		Email:        otp.Email,
		UserPassword: hashedPass,
	})
	if err != nil {
		log.Println(fmt.Errorf("AddNewUser: could not add new user cause %v", err))
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	if err := q.DeleteOtp(r.Context(), user.Email); err != nil {
		log.Printf("adduser: deleteotp: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Printf("cannot load config: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	token, _, err := token.CreateToken(user.UserID, user.Email, "", time.Now().UTC(), config.AccessTokenDuration, false)
	if err != nil {
		log.Println("could not create access token")
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusCreated, token)
}
