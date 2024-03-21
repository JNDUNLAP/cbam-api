package routes

import (
	"dunlap/data"
	"dunlap/files"
	"encoding/json"
	"net/http"
)

func UploadHandler(client *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		report, err := files.XML(r)
		if err != nil {
			errorMsg := "Failed to upload XML file due to an error in processing the request."
			httpStatusCode := http.StatusBadRequest // Use http.StatusBadRequest (400) or http.StatusUnprocessableEntity (422)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpStatusCode)
			json.NewEncoder(w).Encode(map[string]string{"error": errorMsg, "details": err.Error()})

			return
		}

		RespondWithJSON(w, http.StatusOK, report)
	}
}
