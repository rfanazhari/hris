package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// ContractType represents the type of employment contract.
// Allowed values (string representation):
// - "pkwt"
// - "pkwtt"
// - "freelance"
// - "internship"
// Use ParseContractType to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type ContractType string

const (
	ContractPKWT       ContractType = "pkwt"
	ContractPKWTT      ContractType = "pkwtt"
	ContractFreelance  ContractType = "freelance"
	ContractInternship ContractType = "internship"
	ContractPermanent  ContractType = "permanent"
)

func (c ContractType) Valid() bool {
	switch c {
	case ContractPKWT, ContractPKWTT, ContractFreelance, ContractInternship, ContractPermanent:
		return true
	default:
		return false
	}
}

func ParseContractType(s string) (ContractType, error) {
	v := ContractType(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid ContractType: %q", s)
	}
	return v, nil
}

func (c ContractType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

func (c *ContractType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseContractType(s)
	if err != nil {
		return err
	}
	*c = v
	return nil
}

func (c ContractType) Value() (driver.Value, error) {
	if !c.Valid() {
		return nil, fmt.Errorf("invalid ContractType: %q", c)
	}
	return string(c), nil
}

func (c *ContractType) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseContractType(v)
		if err != nil {
			return err
		}
		*c = parsed
		return nil
	case []byte:
		return c.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for ContractType: %T", src)
	}
}
