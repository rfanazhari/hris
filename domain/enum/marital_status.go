package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// MaritalStatus represents a person's marital status.
// Allowed values (string representation):
// - "single"
// - "married"
// - "divorced"
// - "widowed"
// - "separated"
// - "registered_partnership"
// Use ParseMaritalStatus to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type MaritalStatus string

const (
	MaritalSingle                MaritalStatus = "single"
	MaritalMarried               MaritalStatus = "married"
	MaritalDivorced              MaritalStatus = "divorced"
	MaritalWidowed               MaritalStatus = "widowed"
	MaritalSeparated             MaritalStatus = "separated"
	MaritalRegisteredPartnership MaritalStatus = "registered_partnership"
)

func (m MaritalStatus) Valid() bool {
	switch m {
	case MaritalSingle, MaritalMarried, MaritalDivorced, MaritalWidowed, MaritalSeparated, MaritalRegisteredPartnership:
		return true
	default:
		return false
	}
}

func ParseMaritalStatus(s string) (MaritalStatus, error) {
	v := MaritalStatus(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid MaritalStatus: %q", s)
	}
	return v, nil
}

func (m MaritalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(m))
}

func (m *MaritalStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseMaritalStatus(s)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

func (m MaritalStatus) Value() (driver.Value, error) {
	if !m.Valid() {
		return nil, fmt.Errorf("invalid MaritalStatus: %q", m)
	}
	return string(m), nil
}

func (m *MaritalStatus) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseMaritalStatus(v)
		if err != nil {
			return err
		}
		*m = parsed
		return nil
	case []byte:
		return m.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for MaritalStatus: %T", src)
	}
}
