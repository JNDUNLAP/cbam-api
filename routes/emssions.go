package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// GetGoodsEmissions retrieves emissions data for the goods in a specific report.
// swagger:operation GET /reports/{id}/imported-goods/emissions goodsEmissions getGoodsEmissions
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: The ID of the quarterly report to fetch goods emissions details for
//   required: true
//   type: string
// responses:
//   "200":
//     description: Successfully retrieved goods emissions details
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/GoodsEmission'
//   "400":
//     description: Bad request, such as missing report ID, with additional hints provided for resolution
//   "500":
//     description: Internal server error

func GetGoodsEmissions(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		reportID := params["id"]

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
				ErrorDetail: "Failed to fetch the report from the data.",
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
