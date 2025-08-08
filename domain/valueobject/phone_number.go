package valueobject

import (
	"errors"
	"strings"
	"unicode"
)

// PhoneNumber represents a phone number split into country code and local number.
// Example:
//
//	countryCode: "62"
//	number:      "8111020425"
//
// Full() returns the concatenation: "628111020425".
type PhoneNumber struct {
	countryCode string
	number      string
}

// NewPhoneNumber constructs a PhoneNumber after basic normalization and validation.
// - Trims spaces
// - Strips a single leading '+' from countryCode if present
// - Ensures both parts are non-empty and contain only digits
func NewPhoneNumber(countryCode, number string) (*PhoneNumber, error) {
	cc := strings.TrimSpace(countryCode)
	num := strings.TrimSpace(number)

	if strings.HasPrefix(cc, "+") {
		cc = strings.TrimPrefix(cc, "+")
	}

	if cc == "" {
		return nil, errors.New("country code cannot be empty")
	}
	if num == "" {
		return nil, errors.New("number cannot be empty")
	}
	if !isDigits(cc) {
		return nil, errors.New("country code must contain digits only")
	}
	if !isDigits(num) {
		return nil, errors.New("number must contain digits only")
	}

	return &PhoneNumber{countryCode: cc, number: num}, nil
}

// CountryCode returns the country code part (without '+').
func (p PhoneNumber) CountryCode() string { return p.countryCode }

// Number returns the local/national number part as provided (no formatting changes).
func (p PhoneNumber) Number() string { return p.number }

// Full returns the E.164-like concatenation of country code and number without '+'.
// It does not add or remove leading zeros from the number.
func (p PhoneNumber) Full() string { return p.countryCode + p.number }

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
