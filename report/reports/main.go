package reports

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetQuarterlyReport(dbClient *data.MongoDBClient) http.HandlerFunc {
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
			errString := fmt.Sprintf("Failed to fetch the report %s: %s", reportID, err)
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: errString,
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(report); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}

func GetAllQuarterlyReports(dbClient *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		report, err := dbClient.GetQReport(bson.M{})
		if err != nil {
			errString := fmt.Sprintf("Failed to fetch all reports: %s", err)
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: errString,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(report); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
		}
	}
}
func CreateQuarterlyReport(dbClient *data.MongoDBClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collectionName := r.URL.Query().Get("collectionName")
		if collectionName == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Missing collectionName.",
			})
			return
		}

		r.ParseMultipartForm(32 << 20)

		file, _, err := r.FormFile("xmlFile")
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Failed to parse XML file.",
			})
			return
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to read XML file.",
			})
			return
		}

		var newReport model.QReport
		if err := xml.Unmarshal(fileBytes, &newReport); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Failed to unmarshal XML data into QReport struct.",
			})
			return
		}

		err = data.SaveQuarterlyReportToDatabase(&newReport, dbClient, collectionName)
		if err != nil {
			errString := fmt.Sprintf("Failed to save quarterly report: %s", err)
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: errString,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"message": "Quarterly report created successfully"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the response into JSON.",
			})
			return
		}
	}
}

// // createQuarterlyReport creates a new quarterly report with the provided data
// func CreateQuarterlyReport(c echo.Context) error {
// 	// Simulated logic for creating a report; replace with actual data creation logic
// 	report := map[string]interface{}{} // Placeholder for report data
// 	if err := c.Bind(&report); err != nil {
// 		return err
// 	}
// 	// Example data; in a real application, you would insert this into a data
// 	report["id"] = "newly_created_id"
// 	return c.JSON(http.StatusCreated, report)
// }

// // updateQuarterlyReport updates an existing quarterly report identified by ID with new data
// func UpdateQuarterlyReport(c echo.Context) error {
// 	id := c.Param("id")
// 	// Simulated logic for updating a report; replace with actual data update logic
// 	updatedReport := map[string]interface{}{"id": id, "name": "Updated Report Name"} // Example updated data
// 	return c.JSON(http.StatusOK, updatedReport)
// }

// // deleteQuarterlyReport deletes a specific quarterly report by ID
// func DeleteQuarterlyReport(c echo.Context) error {
// 	id := c.Param("id")
// 	fmt.Println(id)
// 	// Simulated logic for deleting a report; replace with actual data deletion logic
// 	// Assuming the report with the specified ID exists and has been deleted successfully
// 	return c.NoContent(http.StatusNoContent)
// }

// // getQuarterlyReports retrieves a list of all quarterly reports
// func GetQuarterlyReports(c echo.Context) error {
// 	// Simulated logic for retrieving reports; replace with actual data retrieval logic
// 	reports := []string{"Report 1", "Report 2", "Report 3"} // Example data
// 	return c.JSON(http.StatusOK, reports)
// }
