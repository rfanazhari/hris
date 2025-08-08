package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestNationality_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.Nationality
		valid bool
	}{
		{"wni valid", enum.NationalityWNI, true},
		{"wna valid", enum.NationalityWNA, true},
		{"invalid value", enum.Nationality("other"), false},
		{"empty value", enum.Nationality(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseNationality(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.Nationality
		wantErr  bool
		testName string
	}{
		{"WNI", enum.NationalityWNI, false, "upper WNI"},
		{" wna ", enum.NationalityWNA, false, "trimmed lower wna"},
		{"x", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseNationality(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for input %q, got nil", tt.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for input %q: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNationality_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	n := enum.NationalityWNI
	b, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"wni\"" {
		t.Fatalf("Marshal got %s, want \"wni\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.Nationality
	if err := json.Unmarshal([]byte("\" WNA \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.NationalityWNA {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.NationalityWNA)
	}

	// Unmarshal invalid
	var u2 enum.Nationality
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid nationality, got nil")
	}
}

func TestNationality_Value(t *testing.T) {
	// Valid value
	v, err := enum.NationalityWNA.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "wna" {
		t.Fatalf("Value() got %#v, want 'wna' string", v)
	}

	// Invalid value
	var invalid enum.Nationality = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestNationality_Scan(t *testing.T) {
	// From string
	var n1 enum.Nationality
	if err := n1.Scan("WNI"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if n1 != enum.NationalityWNI {
		t.Fatalf("Scan(string) got %q, want %q", n1, enum.NationalityWNI)
	}

	// From []byte
	var n2 enum.Nationality
	if err := n2.Scan([]byte("wna")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if n2 != enum.NationalityWNA {
		t.Fatalf("Scan([]byte) got %q, want %q", n2, enum.NationalityWNA)
	}

	// Invalid string value
	var n3 enum.Nationality
	if err := n3.Scan("invalid"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var n4 enum.Nationality
	var src any = 123
	if err := n4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestNationality_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.Nationality
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
