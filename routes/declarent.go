package routes

import (
	"cbam_api/data"
	"cbam_api/errors"
	"cbam_api/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetReportDeclarant(repo data.ReportRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		if r.Method != http.MethodGet {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusMethodNotAllowed,
				Message:     "Method Not Allowed",
				ErrorDetail: "This endpoint only supports GET requests.",
				Hints:       []string{"Use a GET request to access this resource."},
			})
			return
		}

		reportID := params["id"]
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
		report, err := repo.GetQReport(filter)
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
