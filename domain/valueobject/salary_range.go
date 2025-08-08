package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SalaryRange represents a salary band for a role or grade.
// Currency is a 3-letter ISO-like code, currently limited to IDR and USD.
// Values are non-negative and Max must be >= Min.
type SalaryRange struct {
	Min      int64  `json:"min"`
	Max      int64  `json:"max"`
	Currency string `json:"currency"`
}

var allowedCurrencies = map[string]struct{}{
	"IDR": {},
	"USD": {},
}

// NewSalaryRange constructs a SalaryRange with validation and normalization.
func NewSalaryRange(min, max int64, currency string) (*SalaryRange, error) {
	sr := &SalaryRange{Min: min, Max: max, Currency: currency}
	if err := sr.normalizeAndValidate(); err != nil {
		return nil, err
	}
	return sr, nil
}

// Validate checks constraints without normalizing values.
func (s *SalaryRange) Validate() error {
	if s == nil {
		return errors.New("salary range is nil")
	}
	if s.Min < 0 {
		return fmt.Errorf("min must be >= 0")
	}
	if s.Max < s.Min {
		return fmt.Errorf("max must be >= min")
	}
	if s.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}
	if _, ok := allowedCurrencies[s.Currency]; !ok {
		return fmt.Errorf("unsupported currency: %s", s.Currency)
	}
	return nil
}

// normalizeAndValidate uppercases currency then validates.
func (s *SalaryRange) normalizeAndValidate() error {
	if s == nil {
		return errors.New("salary range is nil")
	}
	s.Currency = strings.ToUpper(strings.TrimSpace(s.Currency))
	return s.Validate()
}

// UnmarshalJSON supports lenient currency input (trims and uppercases), then validates.
func (s *SalaryRange) UnmarshalJSON(b []byte) error {
	type alias SalaryRange
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	// copy back and normalize
	s.Min = tmp.Min
	s.Max = tmp.Max
	s.Currency = tmp.Currency
	return s.normalizeAndValidate()
}

// MarshalJSON just marshals the struct as-is.
func (s SalaryRange) MarshalJSON() ([]byte, error) {
	// Ensure normalized before marshaling
	cpy := s
	if err := cpy.normalizeAndValidate(); err != nil {
		return nil, err
	}
	type alias SalaryRange
	return json.Marshal(alias(cpy))
}
