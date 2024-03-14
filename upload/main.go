package upload

import (
	"dunlap/model"
	"encoding/xml"
	"io"
	"net/http"
)

func XML(r *http.Request) (model.QReport, error) {
	r.ParseMultipartForm(32 << 20) // limit your multipart form size

	file, _, err := r.FormFile("xmlFile")
	if err != nil {
		return model.QReport{}, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return model.QReport{}, err
	}

	var report model.QReport
	if err := xml.Unmarshal(fileBytes, &report); err != nil {
		return model.QReport{}, err
	}

	return report, nil
}
