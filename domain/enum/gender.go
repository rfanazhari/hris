package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// Gender represents a person's gender code.
// Allowed values:
//   - "M" (Male)
//   - "F" (Female)
//   - "U" (Unknow)
//
// Use ParseGender to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
	GenderUnknow Gender = "U"
)

func (g Gender) Valid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderUnknow:
		return true
	default:
		return false
	}
}

func ParseGender(s string) (Gender, error) {
	v := Gender(strings.ToUpper(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid Gender: %q", s)
	}
	return v, nil
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(g))
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseGender(s)
	if err != nil {
		return err
	}
	*g = v
	return nil
}

func (g Gender) Value() (driver.Value, error) {
	if !g.Valid() {
		return nil, fmt.Errorf("invalid Gender: %q", g)
	}
	return string(g), nil
}

func (g *Gender) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseGender(v)
		if err != nil {
			return err
		}
		*g = parsed
		return nil
	case []byte:
		return g.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
}
