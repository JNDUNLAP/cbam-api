package remarks

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetRemarksEmissions(dbClient *data.MongoDBClient) http.HandlerFunc {
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
		if err := json.NewEncoder(w).Encode(report.ImportedGoods.Remarks); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}

// // UpdateRemarksEmissions handles updating remarks on emissions for a specific imported good.
// func UpdateRemarksEmissions(c echo.Context) error {
// 	// Implement logic to update remarks on emissions for a good
// 	return c.String(http.StatusOK, "Remarks on emissions updated")
// }
