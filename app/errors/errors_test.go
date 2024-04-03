package errors_test

import (
	"cbam_api/errors"
	"cbam_api/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockError() *model.Error {
	return &model.Error{
		StatusCode:  http.StatusNotFound,
		Message:     "Not Found",
		ErrorDetail: "The requested resource was not found.",
		Hints:       []string{"Check the URL", "Read the docs"},
		Details:     map[string]interface{}{"missing_id": "123"},
	}
}

func TestWriteError(t *testing.T) {

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	testErr := MockError()

	errors.WriteError(w, req, testErr)

	if w.Code != testErr.StatusCode {
		t.Errorf("Expected status code %d, got %d", testErr.StatusCode, w.Code)
	}

	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	expectedBody := `{"message":"Not Found","error":"The requested resource was not found.","hints":["Check the URL","Read the docs"],"details":{"missing_id":"123"}}` + "\n"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
