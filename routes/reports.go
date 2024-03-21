package routes

import (
	"dunlap/data"
	"dunlap/errors"
	"dunlap/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetQuarterlyReport(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		reportID := params["id"]
		log.Println(reportID)
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

func GetAllQuarterlyReports(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {

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

func CreateQuarterlyReport(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
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

func DeleteAllQuarterlyReports(dbClient *data.MongoDBClient) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		if r.Method != http.MethodDelete {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusMethodNotAllowed,
				Message:     "Method Not Allowed",
				ErrorDetail: "This endpoint requires a DELETE request.",
			})
			return
		}

		collectionName := r.URL.Query().Get("collectionName")
		if collectionName == "" {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusBadRequest,
				Message:     "Bad Request",
				ErrorDetail: "Collection name is required as a query parameter.",
			})
			return
		}

		err := dbClient.DeleteAllQReports(collectionName)
		if err != nil {
			errString := fmt.Sprintf("Failed to delete all reports from collection %s: %s", collectionName, err)
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: errString,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "All quarterly reports successfully deleted from the specified collection"}); err != nil {
			errors.WriteError(w, r, &model.Error{
				StatusCode:  http.StatusInternalServerError,
				Message:     "Internal Server Error",
				ErrorDetail: "Failed to encode the success response into JSON.",
			})
		}
	}
}
