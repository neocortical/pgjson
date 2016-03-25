package pgjson

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Postgres' JSONB type. It's a byte array of already encoded JSON (like json.RawMessage)
// which also saves itself correctly to PG's jsonb type.
type JSONB []byte

// NewJSONB creates a new JSONB object by marshaling the supllied object into raw JSON
func NewJSONB(v interface{}) (j JSONB, err error) {
	if v == nil {
		return j, nil
	}

	j, err = json.Marshal(v)
	return
}

func (j JSONB) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

// Scan implements the database/sql.Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	if j == nil {
		return errors.New("Scan called on nil JSONB")
	}

	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		*j = nil
		return errors.New("Scan source was not string")
	}
	// make a copy of the bytes
	*j = append((*j)[0:0], s...)

	return nil
}

// MarshalJSON returns *m as the JSON encoding of m.
func (m JSONB) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *JSONB) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

// Unmarshal unmarshals the raw JSON held by this object into the supplied destination
func (m JSONB) Unmarshal(dest interface{}) (err error) {
	if m == nil {
		return errors.New("trying to unmarshal nil JSONB")
	}
	err = json.Unmarshal([]byte(m), dest)
	return
}

func (j JSONB) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

func (j JSONB) Equals(j1 JSONB) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
