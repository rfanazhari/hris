package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestMaritalStatus_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.MaritalStatus
		valid bool
	}{
		{"single valid", enum.MaritalSingle, true},
		{"married valid", enum.MaritalMarried, true},
		{"divorced valid", enum.MaritalDivorced, true},
		{"widowed valid", enum.MaritalWidowed, true},
		{"separated valid", enum.MaritalSeparated, true},
		{"registered_partnership valid", enum.MaritalRegisteredPartnership, true},
		{"invalid value", enum.MaritalStatus("unknown"), false},
		{"empty value", enum.MaritalStatus(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseMaritalStatus(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.MaritalStatus
		wantErr  bool
		testName string
	}{
		{"single", enum.MaritalSingle, false, "lower single"},
		{" Married ", enum.MaritalMarried, false, "trimmed married"},
		{"DIVORCED", enum.MaritalDivorced, false, "upper divorced"},
		{"WiDoWeD", enum.MaritalWidowed, false, "mixed widowed"},
		{"separated", enum.MaritalSeparated, false, "lower separated"},
		{"REGISTERED_PARTNERSHIP", enum.MaritalRegisteredPartnership, false, "upper registered_partnership"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseMaritalStatus(tt.in)
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

func TestMaritalStatus_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	ms := enum.MaritalRegisteredPartnership
	b, err := json.Marshal(ms)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"registered_partnership\"" {
		t.Fatalf("Marshal got %s, want \"registered_partnership\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.MaritalStatus
	if err := json.Unmarshal([]byte("\" SePaRaTeD \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.MaritalSeparated {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.MaritalSeparated)
	}

	// Unmarshal invalid
	var u2 enum.MaritalStatus
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid marital status, got nil")
	}
}

func TestMaritalStatus_Value(t *testing.T) {
	// Valid value
	v, err := enum.MaritalMarried.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "married" {
		t.Fatalf("Value() got %#v, want 'married' string", v)
	}

	// Invalid value
	var invalid enum.MaritalStatus = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestMaritalStatus_Scan(t *testing.T) {
	// From string
	var m1 enum.MaritalStatus
	if err := m1.Scan("DIVORCED"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if m1 != enum.MaritalDivorced {
		t.Fatalf("Scan(string) got %q, want %q", m1, enum.MaritalDivorced)
	}

	// From []byte
	var m2 enum.MaritalStatus
	if err := m2.Scan([]byte("widowed")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if m2 != enum.MaritalWidowed {
		t.Fatalf("Scan([]byte) got %q, want %q", m2, enum.MaritalWidowed)
	}

	// Invalid string value
	var m3 enum.MaritalStatus
	if err := m3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var m4 enum.MaritalStatus
	var src any = 3.14
	if err := m4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestMaritalStatus_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.MaritalStatus
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
