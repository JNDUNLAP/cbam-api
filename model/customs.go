package model

import (
	"math/big"
	"time"
)

type DateTimeFull struct {
	Required    bool
	Value       time.Time `json:"value"`
	ErrorDetail string    `json:"errorDetail,omitempty"`
}

type SimpleDate struct {
	Required    bool
	Value       string `json:"value"`
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedDecimal struct {
	Value       *big.Float
	Required    bool
	TotalDigits int
	FracDigits  int
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedInt struct {
	Value       int
	Required    bool
	Min         int
	Max         int
	ErrorDetail string `json:"errorDetail,omitempty"`
}

type ConstrainedString struct {
	Value       string
	Required    bool
	Min         int
	Max         int
	ErrorDetail string `json:"errorDetail,omitempty"`
}
