package goods

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetImportedGoods(dbClient *data.MongoDBClient) http.HandlerFunc {
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
		if err := json.NewEncoder(w).Encode(report.ImportedGoods); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}

// // AddImportedGood handles adding a new imported good to a quarterly report.
// func AddImportedGood(c echo.Context) error {
// 	// Implement logic to add a new imported good
// 	return c.String(http.StatusOK, "Imported good added")
// }

// // UpdateImportedGood handles updating an imported good in a quarterly report.
// func UpdateImportedGood(c echo.Context) error {
// 	// Implement logic to update an imported good
// 	return c.String(http.StatusOK, "Imported good updated")
// }

// // DeleteImportedGood handles deleting an imported good from a quarterly report.
// func DeleteImportedGood(c echo.Context) error {
// 	// Implement logic to delete an imported good
// 	return c.String(http.StatusOK, "Imported good deleted")
// }
