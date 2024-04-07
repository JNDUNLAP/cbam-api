package routes

import (
	"cbam_api/model"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ProcessReport(xmlData []byte) (*model.QReport, error) {
	var report model.QReport
	if err := xml.Unmarshal(xmlData, &report); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	if err := model.SetupConstraints(&report); err != nil {
		return nil, fmt.Errorf("failed to apply constraints: %w", err)
	}

	return &report, nil
}

func GenerateJson(xmlData []byte, returnErrors bool) ([]byte, error) {
	xmlFile, err := os.ReadFile("files/xml/Sample_CBAM_Quarterly_Report.xml")
	if err != nil {
		log.Println("Failed to read XML file: %v", err)
	}

	report, err := ProcessReport(xmlFile)
	if err != nil {
		log.Println("Report XML processing failed: %v", err)
	}

	jsonData, err := model.CreateJSON(report, returnErrors)
	if err != nil {
		log.Fatalf("Failed to marshal into JSON: %v", err)
	}

	return jsonData, nil
}

func ExampleReportHandler() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		returnErrors := false

		xmlFile, err := os.ReadFile("files/xml/Sample_CBAM_Quarterly_Report.xml")
		if err != nil {
			log.Printf("Failed to read XML file: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Println("XML file read successfully")

		jsonData, err := GenerateJson(xmlFile, returnErrors)
		if err != nil {
			log.Printf("Report XML processing failed: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonData)
		if err != nil {
			log.Println("Failed to write the JSON response:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
