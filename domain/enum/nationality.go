package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// Nationality represents a person's nationality.
// Allowed values (string representation):
// - "wni"
// - "wna"
// Use ParseNationality to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type Nationality string

const (
	NationalityWNI Nationality = "wni"
	NationalityWNA Nationality = "wna"
)

func (n Nationality) Valid() bool {
	switch n {
	case NationalityWNI, NationalityWNA:
		return true
	default:
		return false
	}
}

func ParseNationality(s string) (Nationality, error) {
	v := Nationality(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid Nationality: %q", s)
	}
	return v, nil
}

func (n Nationality) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(n))
}

func (n *Nationality) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseNationality(s)
	if err != nil {
		return err
	}
	*n = v
	return nil
}

func (n Nationality) Value() (driver.Value, error) {
	if !n.Valid() {
		return nil, fmt.Errorf("invalid Nationality: %q", n)
	}
	return string(n), nil
}

func (n *Nationality) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseNationality(v)
		if err != nil {
			return err
		}
		*n = parsed
		return nil
	case []byte:
		return n.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for Nationality: %T", src)
	}
}
