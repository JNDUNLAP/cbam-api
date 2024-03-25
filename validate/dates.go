package validate

import (
	"fmt"
	"time"
)

func IsValidDateTimeFormat(dt time.Time, layout string) bool {
	if layout != "2006-01-02T15:04:05Z" {
		return false
	}
	dtUTC := dt.UTC()
	formatted := dtUTC.Format(layout)
	parsed, err := time.Parse(layout, formatted)
	if err != nil {
		return false
	}
	isEqual := dtUTC.Truncate(time.Second).Equal(parsed.Truncate(time.Second))
	isUTC := dt.Location() == time.UTC || dt.Location() == time.FixedZone("UTC", 0)
	return isEqual && isUTC
}

func ValidateDateTime(value interface{}) (bool, error) {
	dt, ok := value.(time.Time)
	if !ok {
		return false, fmt.Errorf("expected time.Time, got %T", value)
	}

	if !IsValidDateTimeFormat(dt, "2006-01-02T15:04:05Z") {
		return false, fmt.Errorf("datetime %v does not match the required format '2006-01-02T15:04:05Z' or is not in UTC", dt)
	}

	return true, nil
}
