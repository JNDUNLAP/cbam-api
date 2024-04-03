package model

import (
	"math/big"
	"time"
)

type DateTimeFull struct {
	Value       time.Time `json:"value"`
	ErrorDetail string    `json:"errorDetail,omitempty"`
}

type SimpleDate struct {
	Value       string `json:"value"`
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedDecimal struct {
	Value       *big.Float
	TotalDigits int
	FracDigits  int
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedInt struct {
	Value       int
	Min         int
	Max         int
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedString struct {
	Value       string
	Min         int
	Max         int
	ErrorDetail string `json:"errorDetail,omitempty"`
}
