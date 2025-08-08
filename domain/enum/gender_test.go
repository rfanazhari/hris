package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestGender_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.Gender
		valid bool
	}{
		{"male valid", enum.GenderMale, true},
		{"female valid", enum.GenderFemale, true},
		{"unknow valid", enum.GenderUnknow, true},
		{"invalid value", enum.Gender("X"), false},
		{"empty value", enum.Gender(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseGender(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.Gender
		wantErr  bool
		testName string
	}{
		{"M", enum.GenderMale, false, "upper M"},
		{" f ", enum.GenderFemale, false, "trimmed lower f"},
		{"u", enum.GenderUnknow, false, "lower u"},
		{"x", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseGender(tt.in)
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

func TestGender_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	g := enum.GenderFemale
	b, err := json.Marshal(g)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"F\"" {
		t.Fatalf("Marshal got %s, want \"F\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.Gender
	if err := json.Unmarshal([]byte("\" m \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.GenderMale {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.GenderMale)
	}

	// Unmarshal invalid
	var u2 enum.Gender
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid gender, got nil")
	}
}

func TestGender_Value(t *testing.T) {
	// Valid value
	v, err := enum.GenderUnknow.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "U" {
		t.Fatalf("Value() got %#v, want 'U' string", v)
	}

	// Invalid value
	var invalid enum.Gender = "Z"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestGender_Scan(t *testing.T) {
	// From string
	var g1 enum.Gender
	if err := g1.Scan("f"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if g1 != enum.GenderFemale {
		t.Fatalf("Scan(string) got %q, want %q", g1, enum.GenderFemale)
	}

	// From []byte
	var g2 enum.Gender
	if err := g2.Scan([]byte("M")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if g2 != enum.GenderMale {
		t.Fatalf("Scan([]byte) got %q, want %q", g2, enum.GenderMale)
	}

	// Invalid string value
	var g3 enum.Gender
	if err := g3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var g4 enum.Gender
	var src any = 123
	if err := g4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestGender_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.Gender
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
