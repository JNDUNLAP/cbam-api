package model

import (
	"math/big"
	"time"
)

type DateTimeFull struct {
	MinOcurrences int
	Value         time.Time `json:"value"`
	ErrorDetail   string    `json:"errorDetail,omitempty"`
}

type SimpleDate struct {
	MinOcurrences int
	Value         string `json:"value"`
	ErrorDetail   string `json:"errorDetail,omitempty"`
}

type ConstrainedDecimal struct {
	Value         *big.Float
	MinOcurrences int
	TotalDigits   int
	FracDigits    int
	ErrorDetail   string `json:"errorDetail,omitempty"`
}

type ConstrainedInt struct {
	Value         int
	MinOcurrences int
	Min           int
	Max           int
	ErrorDetail   string `json:"errorDetail,omitempty"`
}

type ConstrainedString struct {
	Value         string
	MinOcurrences int
	Min           int
	Max           int
	ErrorDetail   string `json:"errorDetail,omitempty"`
}
