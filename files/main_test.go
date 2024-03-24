package files_test

import (
	"dunlap/files"
	"os"
	"testing"
)

func TestUnmarshalReport(t *testing.T) {
	const dirPath = "xml/report/"
	const fileName = "Sample_CBAM_Quarterly_Report.xml"

	file, err := os.Open(dirPath + fileName)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	report, err := files.UnmarshalReport(file)
	if err != nil {
		t.Fatalf("UnmarshalReport failed: %v", err)
	}
	if report.DraftReportId == "" {
		t.Errorf("ReportId is empty, but it was expected to be populated")
	}

}
