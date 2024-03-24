package files

import (
	"cbam_api/model"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func XML(r *http.Request) (model.QReport, error) {
	multipartFile, err := extractXMLFileFromRequest(r)
	if err != nil {
		return model.QReport{}, err
	}
	defer multipartFile.Close()

	report, err := UnmarshalReport(multipartFile)
	if err != nil {
		return model.QReport{}, err
	}

	return report, nil
}

func extractXMLFileFromRequest(r *http.Request) (multipartFile io.ReadCloser, err error) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	multipartFile, _, err = r.FormFile("xmlFile")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve 'xmlFile' from form: %w", err)
	}

	return multipartFile, nil
}

func UnmarshalReport(file io.Reader) (model.QReport, error) {
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
