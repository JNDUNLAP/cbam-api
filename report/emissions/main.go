package emissions

import (
	"dunlap/database"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGoodsEmissions(dbClient *database.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reportID := r.URL.Query().Get("reportID")
		if reportID == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing reportID.",
				Hints:       []string{"Ensure 'reportID' is specified in the request parameters."},
			})
			return
		}
		filter := bson.M{"ReportId": reportID}
		report, err := dbClient.GetQReport(filter)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to fetch the report from the database.",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(report.ImportedGoods.GoodsEmissions); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}

// func UpdateGoodsEmissions(c echo.Context) error {
// 	// Implement logic to update emissions data for a good
// 	return c.String(http.StatusOK, "Goods emissions data updated")
// }
