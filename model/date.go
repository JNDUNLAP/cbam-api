package model

import (
	"encoding/xml"
	"strings"
	"time"
)

type SimpleDate struct {
	time.Time
}

func (d *SimpleDate) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var dateStr string
	if err := decoder.DecodeElement(&dateStr, &start); err != nil {
		return err
	}

	// Trim whitespace and check the format
	dateStr = strings.TrimSpace(dateStr)

	// Define a slice of layouts to attempt parsing with.
	// Ensure the date-only layout is included and attempted.
	layouts := []string{
		time.RFC3339,           // "2006-01-02T15:04:05Z07:00" Full datetime with timezone
		"2006-01-02T15:04:05Z", // Full datetime in UTC
		"2006-01-02",           // Date only
	}

	var parseErr error
	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, dateStr)
		if err == nil {
			d.Time = parsedTime
			return nil
		}
		parseErr = err
	}

	// Return the last error if all parse attempts fail.
	return parseErr
}

func (d SimpleDate) FormatISO8601() string {
	return d.Format("2006-01-02T15:04:05Z")
}
