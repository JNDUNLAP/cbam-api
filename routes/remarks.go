package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetRemarksEmissions retrieves remarks and emissions details for imported goods in a specific report.
// swagger:operation GET /reports/{id}/imported-goods/{goodId}/remarks-emissions remarksEmissions getRemarksEmissions
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the quarterly report to fetch remarks and emissions details for
//   required: true
//   type: string
// - name: goodId
//   in: path
//   description: The ID of the imported good within the report
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved remarks and emissions for the specified imported good
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/RemarkEmission'
//   "400":
//     description: Bad request, such as missing report ID or good ID
//   "500":
//     description: Internal server error

func GetRemarksEmissions(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		reportID := params["id"]

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
