package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/big"
)

type ConstrainedDecimal struct {
	Value       *big.Float
	TotalDigits int // Total number of digits allowed
	FracDigits  int // Number of digits allowed after the decimal point
	ErrorDetail string
}

func (c *ConstrainedDecimal) String() string {
	return fmt.Sprintf("%f", c.Value)
}

func (c *ConstrainedDecimal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	parsedValue, _, err := big.ParseFloat(v, 10, 0, big.ToNearestEven)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' as big.Float: %v", v, err)
	}

	c.Value = parsedValue
	return nil
}

func (cs *ConstrainedDecimal) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}
