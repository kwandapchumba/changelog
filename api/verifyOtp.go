package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	sqlc "github.com/kwandapchumba/changelog/db/sqlc"
	"github.com/kwandapchumba/changelog/response"
)

type verifyOtpRequestParams struct {
	Otp string `json:"otp"`
}

func (s verifyOtpRequestParams) validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Otp, validation.Required.Error("otp is required"), validation.Length(6, 6).Error("otp must be six characters long")),
	)
}

func (h *BaseHandler) VerifyOtp(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	var req verifyOtpRequestParams

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
		log.Printf("VerifyOtp: %v: %d", err, http.StatusForbidden)
		response.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	q := sqlc.New(h.db)

	otp, err := q.GetOtp(r.Context(), req.Otp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("VerifyOtp: %v: %d", err, http.StatusNotFound)
			response.Error(w, "otp not found", http.StatusNotFound)
			return
		}

		log.Printf("VerifyOtp: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	if time.Now().UTC().After(otp.Expiry) {
		log.Println("VerifyOtp: otp has expired")
		response.Error(w, "otp has expired", http.StatusForbidden)
		return
	}

	if err := q.UpdateOtp(r.Context(), otp.Otp); err != nil {
		log.Printf("VerifyOtp: UpdateOtp: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusOK, otp.Email)
}
