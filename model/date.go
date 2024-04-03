package model

import (
	"time"
)

func (dt *DateTimeFull) IsValid() bool {
	return !dt.Value.IsZero() && !dt.Value.After(time.Now())
}

func (sd SimpleDate) IsValidFormat() bool {
	_, err := time.Parse("2006-01-02", sd.Value)
	return err == nil
}
