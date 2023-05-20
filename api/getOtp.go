package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	sqlc "github.com/kwandapchumba/changelog/db/sqlc"
	"github.com/kwandapchumba/changelog/response"
	"github.com/kwandapchumba/changelog/utils"
)

type getOtpRequestParams struct {
	Email string `json:"email"`
}

func (h *BaseHandler) GetOtp(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	var req getOtpRequestParams

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

	if req.Email == "" {
		log.Printf("GetOtp: email address is empty: %d", http.StatusForbidden)
		response.Error(w, "email address cannot be empty", http.StatusForbidden)
		return
	}

	q := sqlc.New(h.db)

	params := sqlc.AddOtpParams{
		Otp:    utils.GenerateOTP(),
		Email:  req.Email,
		Expiry: time.Now().UTC().Add(30 * time.Minute),
	}

	otp, err := q.AddOtp(r.Context(), params)
	if err != nil {
		log.Printf("AddOtp: %v: %d", err, http.StatusInternalServerError)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	_ = otp // mail otp to user

	response.JSON(w, http.StatusOK, otp)
}
