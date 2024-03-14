package signatures

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSignatures(dbClient *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reportID := r.URL.Query().Get("reportID")
		if reportID == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing reportID.",
			})
			return
		}
		filter := bson.M{"ReportId": reportID}
		report, err := dbClient.GetQReport(filter)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to fetch the report from the data.",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(report.Signatures); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}

// UpdateSignatures handles updating signatures for a quarterly report.
func UpdateSignatures(c echo.Context) error {
	// Implement logic to update signatures
	return c.String(http.StatusOK, "Signatures updated")
}
