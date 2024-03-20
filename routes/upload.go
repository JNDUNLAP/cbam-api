package routes

import (
	"dunlap/data"
	"dunlap/files"
	"net/http"
)

// UploadHandler uploads a report XML file and returns the processed report.
// swagger:operation POST /upload report uploadHandler
// ---
// consumes:
// - multipart/form-data
// produces:
// - application/json
// parameters:
// - name: xmlFile
//   in: formData
//   description: The XML file containing the report data
//   required: true
//   type: file
// responses:
//   "200":
//     description: Successfully uploaded and processed the report
//     schema:
//       $ref: '#/definitions/Report'
//   "405":
//     description: Method not allowed, this endpoint only supports POST requests
//   "500":
//     description: Internal server error, such as failure in processing the XML file

func UploadHandler(client *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		report, err := files.XML(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RespondWithJSON(w, http.StatusOK, report)
	}
}
