package model

import (
	"encoding/json"
)

type ErrorProvider interface {
	Errors() []string
}

func (r *QReport) Errors() []string {
	return r.ErrorMessages
}

func (r *QReport) AddError(msg string) {
	r.ErrorMessages = append(r.ErrorMessages, msg)
}

func (cs *SimpleDate) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}

func (cs *DateTimeFull) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
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

func (cs *ConstrainedInt) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}

func (cs *ConstrainedString) MarshalJSON() ([]byte, error) {
	if MarshalErrorsOnly {
		if cs.ErrorDetail != "" {
			return json.Marshal(cs.ErrorDetail)
		}
		return json.Marshal(nil)
	}
	return json.Marshal(cs.Value)
}

func CreateJSON(report interface{}, errorsOnly bool) ([]byte, error) {
	MarshalErrorsOnly = errorsOnly
	defer func() { MarshalErrorsOnly = false }()
	return json.MarshalIndent(report, "", "  ")
}
