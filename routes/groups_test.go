package routes_test

import (
	"dunlap/model"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func LoadXMLFromFile(filePath string) (*model.QReport, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var reports QReport
	if err := xml.Unmarshal(byteValue, &reports); err != nil {
		return nil, err
	}

	return &reports, nil
}
