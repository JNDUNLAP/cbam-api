package importer

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetImporter(dbClient *data.MongoDBClient) http.HandlerFunc {
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

// // UpdateImporter updates the importer information for a specific quarterly report
// func UpdateImporter(c echo.Context) error {
// 	id := c.Param("id")
// 	// Simulated logic for updating importer info; replace with actual data update logic
// 	updatedImporter := map[string]interface{}{} // Placeholder for importer data
// 	if err := c.Bind(&updatedImporter); err != nil {
// 		return err
// 	}
// 	// Example data; in a real application, you would update this in a data
// 	updatedImporter["reportID"] = id
// 	return c.JSON(http.StatusOK, updatedImporter)
// }
