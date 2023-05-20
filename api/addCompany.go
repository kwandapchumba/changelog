package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	sqlc "github.com/kwandapchumba/changelog/db/sqlc"
	middleware "github.com/kwandapchumba/changelog/middleware"
	"github.com/kwandapchumba/changelog/response"
	"github.com/kwandapchumba/changelog/token"
	"github.com/kwandapchumba/changelog/utils"
)

type addCompanyRequestParams struct {
	CompanyName string `json:"company_name"`
}

func (h *BaseHandler) AddCompany(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)

	d.DisallowUnknownFields()

	var req addCompanyRequestParams

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

	if req.CompanyName == "" {
		log.Println("AddCompany: company name empty")
		response.Error(w, "company name is required", http.StatusForbidden)
		return
	}

	const pLoad middleware.ContextKey = "payload"

	payload := r.Context().Value(pLoad).(*token.Payload)

	q := sqlc.New(h.db)

	company, err := q.AddCompany(r.Context(), sqlc.AddCompanyParams{
		CompanyID:   utils.RandomString(),
		CompanyName: req.CompanyName,
		UserID:      payload.UserID,
	})
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "company_pkey" (SQLSTATE 23505)` {
			msg := "user has an existing company"
			log.Printf("AddCompany: %s: %d", msg, http.StatusConflict)
			response.Error(w, msg, http.StatusConflict)
			return
		}

		log.Printf("AddCompany: %v", err.Error())
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Printf("cannot load config: %v", err)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	token, _, err := token.CreateToken(payload.UserID, payload.Email, company.CompanyID, time.Now().UTC(), config.AccessTokenDuration, true)
	if err != nil {
		log.Printf("AddCompany: could not create access token: %d", http.StatusInternalServerError)
		response.Error(w, internalError, http.StatusInternalServerError)
		return
	}

	response.JSON(w, http.StatusCreated, token)
}
