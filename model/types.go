package model

import (
	"regexp"
)

var (
	PatternAlphaNumeric            = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	PatternNumeric                 = regexp.MustCompile(`^\d+$`)
	PatternAlphaNumericWithSpaces  = regexp.MustCompile(`^[a-zA-Z0-9\s]*$`)
	PatternAlphaNumericSpecialChar = regexp.MustCompile(`^[a-zA-Z0-9\s,.'-]*$`)
	PatternISODateTime             = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
)
