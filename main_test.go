package main_test

import (
	"cbam_api/validate"
	"testing"
)

func TestDataModelValidate(t *testing.T) {
	if err := validate.DataModelValidate(); err != nil {
		t.Fatalf("The Example File \n(files/xml/report/Sample_CBAM_Quarterly_Report.xml)\nfailed to data model valiation: %v", err)
	}
}
