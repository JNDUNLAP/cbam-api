package routes

import (
	"cbam_api/data"
	"cbam_api/model"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func UploadHandler(client *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var report model.QReport
		err := xml.NewDecoder(r.Body).Decode(&report)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Failed to decode XML: "+err.Error())
			return
		}

		// errorPayload, err := model.CreateJSON(report, false)
		// if err != nil {
		// 	respondWithError(w, http.StatusInternalServerError, "Failed to convert report to Errors to JSON: "+err.Error())
		// 	return
		// }

		jsonData, err := model.CreateJSON(&report, true)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to convert report to JSON: "+err.Error())
			return
		}

		if err := data.UploadReport(r.Context(), &report, client, "test_reports"); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to upload report: "+err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
