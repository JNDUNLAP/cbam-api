package model

import (
	"encoding/json"
	"encoding/xml"
)

type ConstrainedString struct {
	Value       string
	Min         int
	Max         int
	ErrorDetail string
}

func (c *ConstrainedString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	c.Value = v
	return nil
}

func (cs *ConstrainedString) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}
