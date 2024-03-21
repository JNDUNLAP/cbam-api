package files

import (
	"dunlap/model"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func XML(r *http.Request) (model.QReport, error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return model.QReport{}, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	file, _, err := r.FormFile("xmlFile")
	if err != nil {
		return model.QReport{}, fmt.Errorf("failed to retrieve 'xmlFile' from form: %w", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return model.QReport{}, fmt.Errorf("failed to read file contents: %w", err)
	}

	var report model.QReport
	if err := xml.Unmarshal(fileBytes, &report); err != nil {
		return model.QReport{}, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	return report, nil
}
