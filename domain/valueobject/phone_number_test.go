package valueobject_test

import (
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewPhoneNumber_Valid(t *testing.T) {
	p, err := vo.NewPhoneNumber("62", "8111020425")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.CountryCode() != "62" {
		t.Fatalf("CountryCode mismatch: got %s", p.CountryCode())
	}
	if p.Number() != "8111020425" {
		t.Fatalf("Number mismatch: got %s", p.Number())
	}
	if p.Full() != "628111020425" {
		t.Fatalf("Full() mismatch: got %s, want %s", p.Full(), "628111020425")
	}
}

func TestNewPhoneNumber_NormalizesPlus(t *testing.T) {
	p, err := vo.NewPhoneNumber("+62", "8111")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.CountryCode() != "62" {
		t.Fatalf("expected '+' to be stripped, got %s", p.CountryCode())
	}
}

func TestNewPhoneNumber_LeadingZeroPreserved(t *testing.T) {
	p, err := vo.NewPhoneNumber("62", "0811")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Full() != "620811" {
		t.Fatalf("Full() should preserve leading zero: got %s", p.Full())
	}
}

func TestNewPhoneNumber_InvalidInputs(t *testing.T) {
	cases := []struct {
		name string
		cc   string
		num  string
	}{
		{"empty cc", "", "123"},
		{"empty num", "62", ""},
		{"non-digit cc", "6a", "123"},
		{"non-digit num", "62", "12b"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewPhoneNumber(c.cc, c.num); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
