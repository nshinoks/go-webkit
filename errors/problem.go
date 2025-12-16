package errors

import (
	"encoding/json"
	"net/http"
)

type Problem struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	// Extensions: allow non-standard fields without breaking clients
	Extensions map[string]any `json:"-"`
}

func (p Problem) With(key string, value any) Problem {
	if p.Extensions == nil {
		p.Extensions = map[string]any{}
	}
	p.Extensions[key] = value
	return p
}

func (p Problem) MarshalJSON() ([]byte, error) {
	type alias Problem
	a := map[string]any{}
	b, err := json.Marshal(alias(p))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return nil, err
	}
	for k, v := range p.Extensions {
		a[k] = v
	}
	return json.Marshal(a)
}

// Write writes RFC7807-like JSON (problem+json)
func (p Problem) Write(w http.ResponseWriter) {
	status := p.Status
	if status == 0 {
		status = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(p)
}
