package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
)

type ConstrainedInt struct {
	Value       int
	Min         int
	Max         int
	ErrorDetail string
}

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

func (cs *ConstrainedInt) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}
