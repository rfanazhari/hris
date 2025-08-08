package valueobject_test

import (
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewEmailAddress_Valid(t *testing.T) {
	e, err := vo.NewEmailAddress("arfanazh", "gmail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if e.Username() != "arfanazh" {
		t.Fatalf("Username mismatch: got %s", e.Username())
	}
	if e.Domain() != "gmail.com" {
		t.Fatalf("Domain mismatch: got %s", e.Domain())
	}
	if e.Full() != "arfanazh@gmail.com" {
		t.Fatalf("Full() mismatch: got %s, want %s", e.Full(), "arfanazh@gmail.com")
	}
}

func TestNewEmailAddress_Normalization(t *testing.T) {
	e, err := vo.NewEmailAddress(" User ", " GMAIL.COM ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// username is trimmed but case-preserved
	if e.Username() != "User" {
		t.Fatalf("expected trimmed username 'User', got %q", e.Username())
	}
	// domain is trimmed and lowercased
	if e.Domain() != "gmail.com" {
		t.Fatalf("expected lowercased domain 'gmail.com', got %q", e.Domain())
	}
	if e.Full() != "User@gmail.com" {
		t.Fatalf("Full() mismatch: got %s, want %s", e.Full(), "User@gmail.com")
	}
}

func TestNewEmailAddress_InvalidInputs(t *testing.T) {
	cases := []struct {
		name     string
		username string
		domain   string
	}{
		{"empty username", "", "gmail.com"},
		{"empty domain", "user", ""},
		{"username has space", "ar fan", "gmail.com"},
		{"username has @", "ar@fan", "gmail.com"},
		{"domain has space", "user", "gmail .com"},
		{"domain has @", "user", "gma@il.com"},
		{"domain starts with dot", "user", ".gmail.com"},
		{"domain ends with dot", "user", "gmail.com."},
		{"domain without dot", "user", "gmail"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewEmailAddress(c.username, c.domain); err == nil {
				t.Fatalf("expected error, got nil for %q@%q", c.username, c.domain)
			}
		})
	}
}
