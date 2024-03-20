package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetReportDeclarant retrieves declarant information for a specific report.
// swagger:operation GET /reports/{id}/declarant declarant getReportDeclarant
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the report for which to retrieve declarant information
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved declarant information
//     schema:
//       $ref: '#/definitions/Declarant'
//   "400":
//     description: Bad request, such as missing report ID, with hints for resolution
//   "405":
//     description: Method not allowed, this endpoint only supports GET requests
//   "500":
//     description: Internal server error

func GetReportDeclarant(repo data.ReportRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {

		if r.Method != http.MethodGet {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusMethodNotAllowed,
				Message:     "Method Not Allowed",
				ErrorDetail: "This endpoint only supports GET requests.",
				Hints:       []string{"Use a GET request to access this resource."},
			})
			return
		}

		reportID := params["id"]
		if reportID == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing report ID.",
				Hints:       []string{"Include a 'reportID' query parameter with your request."},
			})
			return
		}

		filter := bson.M{"reportId": reportID}
		report, err := repo.GetQReport(filter)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to fetch report.",
				Hints:       []string{""},
			})
			return
		}

		if err := json.NewEncoder(w).Encode(report.Declarant); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode response.",
				Hints:       []string{""},
			})
			return
		}
	}
}
