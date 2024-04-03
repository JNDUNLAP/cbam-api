package errors

import (
	"cbam_api/model"
	"encoding/json"
	"log"
	"net/http"
)

func LogError(r *http.Request, err *model.Error) {
	log.Printf("Error %d - %s: %s. Hints: %v, Details: %v, Requested by: %s",
		err.StatusCode, err.Message, err.ErrorDetail, err.Hints, err.Details, r.RemoteAddr)

}

func WriteError(w http.ResponseWriter, req *http.Request, err *model.Error) {
	LogError(req, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Printf("Failed to encode and send error: %v", err)
	}
}
