package valueobject_test

import (
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewEmployeeName_ValidWithMiddle(t *testing.T) {
	name, err := vo.NewEmployeeName("John", "Ronald", "Reuel", "JRR")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := name.FirstName(); got != "John" {
		t.Fatalf("FirstName mismatch: got %s", got)
	}
	if got := name.MiddleName(); got != "Ronald" {
		t.Fatalf("MiddleName mismatch: got %s", got)
	}
	if got := name.LastName(); got != "Reuel" {
		t.Fatalf("LastName mismatch: got %s", got)
	}
	if got := name.NickName(); got != "JRR" {
		t.Fatalf("NickName mismatch: got %s", got)
	}
	if got := name.FullName(); got != "John Ronald Reuel" {
		t.Fatalf("FullName mismatch: got %q, want %q", got, "John Ronald Reuel")
	}
}

func TestNewEmployeeName_ValidWithoutMiddle(t *testing.T) {
	name, err := vo.NewEmployeeName("Jane", "", "Doe", "Janey")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := name.FullName(); got != "Jane Doe" {
		t.Fatalf("FullName mismatch: got %q, want %q", got, "Jane Doe")
	}
}

func TestNewEmployeeName_TrimsSpacesAndSkipsEmptyMiddle(t *testing.T) {
	name, err := vo.NewEmployeeName("  Alice  ", "   ", "  Smith ", "  Ally ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := name.FirstName(); got != "Alice" {
		t.Fatalf("expected trimmed first name, got %q", got)
	}
	if got := name.MiddleName(); got != "" {
		t.Fatalf("expected empty middle after trimming spaces, got %q", got)
	}
	if got := name.LastName(); got != "Smith" {
		t.Fatalf("expected trimmed last name, got %q", got)
	}
	if got := name.NickName(); got != "Ally" {
		t.Fatalf("expected trimmed nickname, got %q", got)
	}
	if got := name.FullName(); got != "Alice Smith" {
		t.Fatalf("FullName should not have double spaces: got %q", got)
	}
}

func TestNewEmployeeName_Invalid(t *testing.T) {
	cases := []struct {
		name string
		fn   string
		mn   string
		ln   string
		nn   string
	}{
		{"empty first", "", "m", "l", "n"},
		{"empty last", "f", "m", "", "n"},
		{"first only spaces", "   ", "m", "l", "n"},
		{"last only spaces", "f", "m", "   ", "n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewEmployeeName(c.fn, c.mn, c.ln, c.nn); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}
