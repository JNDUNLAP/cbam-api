package model

import (
	"encoding/xml"
	"fmt"
	"math/big"
)

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
