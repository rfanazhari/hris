package valueobject_test

import (
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewAddress_Valid(t *testing.T) {
	addr, err := vo.NewAddress("  Jl. Merdeka 10 ", " Jakarta ", " DKI ", " 10110 ", " Indonesia ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := addr.Street(); got != "Jl. Merdeka 10" {
		t.Fatalf("Street mismatch: got %q", got)
	}
	if got := addr.City(); got != "Jakarta" {
		t.Fatalf("City mismatch: got %q", got)
	}
	if got := addr.State(); got != "DKI" {
		t.Fatalf("State mismatch: got %q", got)
	}
	if got := addr.PostalCode(); got != "10110" {
		t.Fatalf("PostalCode mismatch: got %q", got)
	}
	if got := addr.Country(); got != "Indonesia" {
		t.Fatalf("Country mismatch: got %q", got)
	}
}

func TestNewAddress_InvalidInputs(t *testing.T) {
	cases := []struct {
		name       string
		street     string
		city       string
		state      string
		postalCode string
		country    string
	}{
		{"empty street", "", "City", "State", "12345", "Country"},
		{"empty city", "Street", "", "State", "12345", "Country"},
		{"empty state", "Street", "City", "", "12345", "Country"},
		{"empty postal code", "Street", "City", "State", "", "Country"},
		{"empty country", "Street", "City", "State", "12345", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewAddress(c.street, c.city, c.state, c.postalCode, c.country); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
