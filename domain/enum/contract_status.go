package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// ContractStatus represents the status of a contract.
// Allowed values (string representation):
// - "active"
// - "expired"
// - "terminated"
// Use ParseContractStatus to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type ContractStatus string

const (
	ContractStatusActive     ContractStatus = "active"
	ContractStatusExpired    ContractStatus = "expired"
	ContractStatusTerminated ContractStatus = "terminated"
)

func (c ContractStatus) Valid() bool {
	switch c {
	case ContractStatusActive, ContractStatusExpired, ContractStatusTerminated:
		return true
	default:
		return false
	}
}

func ParseContractStatus(s string) (ContractStatus, error) {
	v := ContractStatus(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid ContractStatus: %q", s)
	}
	return v, nil
}

func (c ContractStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

func (c *ContractStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseContractStatus(s)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

func (c ContractStatus) Value() (driver.Value, error) {
	if !c.Valid() {
		return nil, fmt.Errorf("invalid ContractStatus: %q", c)
	}
	return string(c), nil
}

func (c *ContractStatus) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseContractStatus(v)
		if err != nil {
			return err
		}
		*c = parsed
		return nil
	case []byte:
		return c.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for ContractStatus: %T", src)
	}
}
