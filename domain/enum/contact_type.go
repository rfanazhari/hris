package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// ContactType represents type of a contact method for a person/employee.
// Allowed values: "primary", "emergency", "secondary", "work".
// Use ParseContactType to safely convert from string.
// Implements json (un)marshaling and database/sql interfaces.
type ContactType string

const (
	ContactPrimary   ContactType = "primary"
	ContactEmergency ContactType = "emergency"
	ContactSecondary ContactType = "secondary"
	ContactWork      ContactType = "work"
)

func (c ContactType) Valid() bool {
	switch c {
	case ContactPrimary, ContactEmergency, ContactSecondary, ContactWork:
		return true
	default:
		return false
	}
}

func ParseContactType(s string) (ContactType, error) {
	v := ContactType(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid ContactType: %q", s)
	}
	return v, nil
}

func (c ContactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

func (c *ContactType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseContactType(s)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

func (c ContactType) Value() (driver.Value, error) {
	if !c.Valid() {
		return nil, fmt.Errorf("invalid ContactType: %q", c)
	}
	return string(c), nil
}

func (c *ContactType) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseContactType(v)
		if err != nil {
			return err
		}
		*c = parsed
		return nil
	case []byte:
		return c.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for ContactType: %T", src)
	}
}
