package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetNationalCompetentAuth retrieves national competent authority information for a specific report.
// swagger:operation GET /reports/{id}/national-competent-auth nationalCompetentAuth getNationalCompetentAuth
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the quarterly report to fetch national competent authority information for
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved national competent authority information
//     schema:
//       $ref: '#/definitions/NationalCompetentAuthority'
//   "400":
//     description: Bad request, such as missing report ID
//   "500":
//     description: Internal server error

func GetNationalCompetentAuth(dbClient *data.MongoDBClient) HandlerFunc {
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
		if err := json.NewEncoder(w).Encode(report.NationalCompetentAuth); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}
