package main_test

import (
	"cbam_api/model"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnmarshalQReport(t *testing.T) {

	xmlFile, err := ioutil.ReadFile("files/xml/Sample_CBAM_Quarterly_Report.xml")
	if err != nil {
		t.Fatalf("Failed to read XML file: %v", err)
	}
	report, err := ProcessReport(xmlFile)
	if err != nil {
		t.Fatalf("Report XML failed: %v", err)
	}
	// For errors only
	errorPayload, err := CreateJSON(report, true)
	if err != nil {
		fmt.Println("Error creating error-only JSON payload:", err)
		return
	}
	fmt.Println("Error Payload:", string(errorPayload))

}

func ProcessReport(xmlData []byte) (*model.QReport, error) {
	var report model.QReport

	if err := xml.Unmarshal(xmlData, &report); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	if err := model.SetupConstraints(&report); err != nil {
		return nil, fmt.Errorf("failed to set constraints on XML: %w", err)
	}

	// model.PrintStructValues(report, "")
	return &report, nil
}

func CreateJSON(report *model.QReport, errorsOnly bool) ([]byte, error) {
	// Set the flag based on the function call context
	model.MarshalErrorsOnly = errorsOnly
	defer func() { model.MarshalErrorsOnly = false }() // Ensure we reset the flag after use

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
