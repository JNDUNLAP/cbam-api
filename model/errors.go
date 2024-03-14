package model

type Error struct {
	StatusCode  int         `json:"-"`
	Message     string      `json:"message"`
	ErrorDetail string      `json:"error,omitempty"`
	Hints       []string    `json:"hints,omitempty"`
	Details     interface{} `json:"details,omitempty"`
}
