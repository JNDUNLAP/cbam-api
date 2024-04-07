package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

func (c *ConstrainedInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	intVal, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("ConstrainedInt: %v", err)
	}

	c.Value = intVal
	return nil
}
