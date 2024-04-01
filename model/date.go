package model

import (
	"encoding/xml"
	"time"
)

type DateTimeFull struct {
	time.Time
}

func (sd *DateTimeFull) IsValid() bool {
	if sd.IsZero() {
		return false
	}

	if sd.After(time.Now()) {
		return false
	}

	return true
}

type SimpleDate string

func (sd SimpleDate) IsValidFormat() bool {
	_, err := time.Parse("2006-01-02", string(sd))
	return err == nil
}

func GenerateUnmarshalXMLDate(typeName string) func(*xml.Decoder, xml.StartElement) error {
	return func(d *xml.Decoder, start xml.StartElement) error {
		var v SimpleDate
		if err := d.DecodeElement(&v, &start); err != nil {
			return err
		}
		// Validate the date if needed
		return nil
	}
}
