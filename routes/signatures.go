package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetSignatures retrieves signatures associated with a specific quarterly report by its ID.
// swagger:operation GET /reports/{id}/signatures signatures getSignatures
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the quarterly report for which to fetch signatures
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved signatures for the quarterly report
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/Signature'
//   "400":
//     description: Bad request, such as missing report ID
//   "500":
//     description: Internal server error

func GetSignatures(dbClient *data.MongoDBClient) HandlerFunc {
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
		if err := json.NewEncoder(w).Encode(report.Signatures); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}
