package main_test

import (
	"cbam_api/model"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestUnmarshalQReport(t *testing.T) {
	xmlFile, err := ioutil.ReadFile("files/xml/Sample_CBAM_Quarterly_Report.xml")
	if err != nil {
		t.Fatalf("Failed to read XML file: %v", err)
	}

	report, err := ProcessReport(xmlFile)
	if err != nil {
		t.Fatalf("Report XML processing failed: %v", err)
	}

	jsonData, err := model.CreateJSON(report, true)
	if err != nil {
		log.Fatalf("Failed to marshal into JSON: %v", err)
	}

	fmt.Printf("JSON output:\n%s\n", string(jsonData))
}

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
