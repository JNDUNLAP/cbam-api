package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetSupportingDocuments retrieves supporting documents for a specific good in a report.
// swagger:operation GET /reports/{id}/imported-goods/{goodId}/supporting-documents supportingDocuments getSupportingDocuments
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the quarterly report
//   required: true
//   type: string
// - name: goodId
//   in: path
//   description: The ID of the imported good within the report
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved supporting documents for the specified good
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/SupportingDocument'
//   "400":
//     description: Bad request, such as missing report ID or good ID, with additional hints for resolution
//   "500":
//     description: Internal server error

func GetSupportingDocuments(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		reportID := params["id"]
		goodID := params["goodId"]

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
