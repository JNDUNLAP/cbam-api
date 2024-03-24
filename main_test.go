package main_test

import (
	"dunlap/model"
	"testing"
)

func TestDataModelValidate(t *testing.T) {
	if err := model.DataModelValidate(); err != nil {
		t.Fatalf("DataModelTest failed: %v", err)
	}
}
