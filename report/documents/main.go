package documents

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"

	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSupportingDocuments(dbClient *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reportID := r.URL.Query().Get("reportID")
		goodID := r.URL.Query().Get("goodID")

		if reportID == "" || goodID == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing reportID or goodID.",
				Hints:       []string{"Ensure both 'reportID' and 'goodID' are specified in the request parameters."},
			})
			return
		}

		filter := bson.M{"ReportId": reportID, "ImportedGood.GoodId": goodID}
		report, err := dbClient.GetQReport(filter)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to fetch the report from the data.",
				Hints:       []string{},
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(report.ImportedGoods.SupportingDocuments); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
				Hints:       []string{},
			})
		}
	}
}
