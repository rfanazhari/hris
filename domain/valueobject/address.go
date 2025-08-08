package valueobject

import (
	"errors"
	"strings"
)

// Address represents a postal address broken down into common components.
// All fields are stored trimmed. No locale-specific validation is performed
// beyond ensuring fields are non-empty.
//
// Fields:
//   - street
//   - city
//   - state
//   - postalCode
//   - country
//
// Adjust or extend validation externally if stronger guarantees are needed.
type Address struct {
	street     string
	city       string
	state      string
	postalCode string
	country    string
}

// NewAddress constructs an Address with basic normalization and validation.
// - Trims spaces on all fields
// - Ensures all fields are non-empty
func NewAddress(street, city, state, postalCode, country string) (*Address, error) {
	s := strings.TrimSpace(street)
	c := strings.TrimSpace(city)
	st := strings.TrimSpace(state)
	pc := strings.TrimSpace(postalCode)
	co := strings.TrimSpace(country)

	if s == "" {
		return nil, errors.New("street cannot be empty")
	}
	if c == "" {
		return nil, errors.New("city cannot be empty")
	}
	if st == "" {
		return nil, errors.New("state cannot be empty")
	}
	if pc == "" {
		return nil, errors.New("postal code cannot be empty")
	}
	if co == "" {
		return nil, errors.New("country cannot be empty")
	}

	return &Address{
		street:     s,
		city:       c,
		state:      st,
		postalCode: pc,
		country:    co,
	}, nil
}

// Street returns the street line of the address.
func (a Address) Street() string { return a.street }

// City returns the city/locality of the address.
func (a Address) City() string { return a.city }

// State returns the state/province/region of the address.
func (a Address) State() string { return a.state }

// PostalCode returns the postal/ZIP code of the address.
func (a Address) PostalCode() string { return a.postalCode }

// Country returns the country of the address.
func (a Address) Country() string { return a.country }
