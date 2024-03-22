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
	dateStr = strings.TrimSpace(dateStr)
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z",
		"2006-01-02",
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

	return parseErr
}

func (d SimpleDate) FormatISO8601() string {
	return d.Format("2006-01-02T15:04:05Z")
}
