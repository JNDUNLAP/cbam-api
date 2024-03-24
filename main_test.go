package main_test

import (
	"cbam_api/model"
	"testing"
)

func TestDataModelValidate(t *testing.T) {
	if err := model.DataModelValidate(); err != nil {
		t.Fatalf("The Example File \n(files/xml/report/Sample_CBAM_Quarterly_Report.xml)\nfailed to data model valiation: %v", err)
	}
}
