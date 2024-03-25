package validate_test

import (
	"cbam_api/validate"
	"testing"
	"time"
)

func TestIsValidDateTimeFormat(t *testing.T) {
	tests := []struct {
		dt     time.Time
		layout string
		want   bool
	}{
		{time.Date(2023, 8, 8, 8, 21, 14, 0, time.UTC), "2006-01-02T15:04:05Z", true},
		{time.Date(2023, 8, 8, 9, 21, 14, 0, time.FixedZone("CET", 1*3600)), "2006-01-02T15:04:05Z", false},
		{time.Date(2023, 8, 8, 7, 21, 14, 0, time.FixedZone("EST", -5*3600)), "2006-01-02T15:04:05Z", false},
		{time.Date(2023, 8, 8, 8, 21, 14, 500, time.UTC), "2006-01-02T15:04:05Z", true},
		{time.Date(2023, 8, 8, 8, 21, 14, 0, time.FixedZone("PST", -8*3600)), "2006-01-02T15:04:05Z", false},
	}

	for _, tt := range tests {
		got := validate.IsValidDateTimeFormat(tt.dt, tt.layout)
		if got != tt.want {
			t.Errorf("IsValidDateTimeFormat(%v, %q) = %v; want %v", tt.dt, tt.layout, got, tt.want)
		}
	}
}
