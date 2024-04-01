package validate

import (
	"fmt"
	"time"
)

func IsValidDateTimeFormat(dt time.Time, layout string) bool {
	parsed, err := time.Parse(time.RFC3339, dt.Format(time.RFC3339))
	if err != nil {
		return false
	}

	isEqual := dt.Equal(parsed)
	isUTC := dt.Location() == time.UTC || dt.Location() == time.FixedZone("UTC", 0)
	return isEqual && isUTC
}

func ValidateDateTime(value interface{}, typeName string) (bool, error) {
	fmt.Println("Validating datetime:", typeName) // Debugging
	dt, ok := value.(time.Time)
	if !ok {
		return false, fmt.Errorf("expected time.Time, got %T", value)
	}

	if !IsValidDateTimeFormat(dt, "2006-01-02T15:04:05Z") {
		return false, fmt.Errorf("datetime %v does not match the required format '2006-01-02T15:04:05Z' or is not in UTC", dt)
	}

	return true, nil
}
