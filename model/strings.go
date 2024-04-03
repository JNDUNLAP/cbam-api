package model

import (
	"encoding/xml"
)

func (c *ConstrainedString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	c.Value = v
	return nil
}
