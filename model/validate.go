package model

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *ConstrainedDecimal) ValidateConstraints() error {
	if c.Value == nil {
		return fmt.Errorf("null")
	}

	valStr := c.Value.Text('f', -1)
	parts := strings.Split(valStr, ".")
	totalDigits := len(parts[0])
	fracDigits := 0
	if len(parts) > 1 {
		fracDigits = len(parts[1])
	}

	if totalDigits > c.TotalDigits || fracDigits > c.FracDigits {
		c.ErrorDetail = fmt.Sprintf("value %s exceeds allowed digits (total: %d, fractional: %d)", valStr, c.TotalDigits, c.FracDigits)
		return fmt.Errorf(c.ErrorDetail)
	}

	return nil
}

func (c *ConstrainedDecimal) ApplyConstraints(constraints map[string]string) {
	if totalDigits, ok := constraints["totalDigits"]; ok {
		if td, err := strconv.Atoi(totalDigits); err == nil {
			c.TotalDigits = td
		}
	}
	if fracDigits, ok := constraints["fracDigits"]; ok {
		if fd, err := strconv.Atoi(fracDigits); err == nil {
			c.FracDigits = fd
		}
	}
}

func (c *ConstrainedString) ValidateConstraints() error {
	if c.Min != 0 || c.Max != 0 {
		valueLength := len(c.Value)
		if valueLength < c.Min || valueLength > c.Max {
			err := fmt.Errorf("ERROR: length %d is out of bounds [%d, %d] for ConstrainedString", valueLength, c.Min, c.Max)
			c.ErrorDetail = err.Error()
			fmt.Println(c.ErrorDetail)
			return err
		}
	}
	fmt.Println("ConstrainedString validation passed")
	c.ErrorDetail = ""
	return nil
}

func (c *ConstrainedInt) ValidateConstraints() error {
	valueStr := strconv.Itoa(c.Value)
	valueLength := len(valueStr)
	if valueLength < c.Min || valueLength > c.Max {
		err := fmt.Errorf("ERROR: length %d is out of bounds [%d, %d]", valueLength, c.Min, c.Max)
		c.ErrorDetail = err.Error()
		return err
	}
	c.ErrorDetail = ""
	return nil
}
