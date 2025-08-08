package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// EmploymentStatus represents the status of an employment/employee.
// Allowed values (string representation):
// - "active"
// - "resigned"
// - "on_leave"
// Use ParseEmploymentStatus to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type EmploymentStatus string

const (
	EmploymentActive   EmploymentStatus = "active"
	EmploymentResigned EmploymentStatus = "resigned"
	EmploymentOnLeave  EmploymentStatus = "on_leave"
)

func (e EmploymentStatus) Valid() bool {
	switch e {
	case EmploymentActive, EmploymentResigned, EmploymentOnLeave:
		return true
	default:
		return false
	}
}

func ParseEmploymentStatus(s string) (EmploymentStatus, error) {
	v := EmploymentStatus(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid EmploymentStatus: %q", s)
	}
	return v, nil
}

func (e EmploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(e))
}

func (e *EmploymentStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseEmploymentStatus(s)
	if err != nil {
		return err
	}
	*e = v
	return nil
}

func (e EmploymentStatus) Value() (driver.Value, error) {
	if !e.Valid() {
		return nil, fmt.Errorf("invalid EmploymentStatus: %q", e)
	}
	return string(e), nil
}

func (e *EmploymentStatus) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseEmploymentStatus(v)
		if err != nil {
			return err
		}
		*e = parsed
		return nil
	case []byte:
		return e.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for EmploymentStatus: %T", src)
	}
}
