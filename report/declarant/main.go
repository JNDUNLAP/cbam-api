package declarant

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetReportDeclarant(client *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusMethodNotAllowed,
				Message:     "Method Not Allowed",
				ErrorDetail: "This endpoint only supports GET requests.",
				Hints:       []string{"Use a GET request to access this resource."},
			})
			return
		}

		reportID := r.URL.Query().Get("reportID")
		if reportID == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing report ID.",
				Hints:       []string{"Include a 'reportID' query parameter with your request."},
			})
			return
		}

		filter := bson.M{"reportId": reportID}
		report, err := client.GetQReport(filter)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to fetch report.",
				Hints:       []string{""},
			})
			return
		}

		if err := json.NewEncoder(w).Encode(report.Declarant); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode response.",
				Hints:       []string{""},
			})
			return
		}
	}
}
