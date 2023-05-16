package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(httpStatusCode)

	resp := make(map[string]string)

	resp["message"] = message

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Printf("failed to marshal reponse at response.go, ErrorResponse function: %v", err)
		return
	}

	w.Write(jsonResponse)
}
