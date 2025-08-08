package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestContactType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.ContactType
		valid bool
	}{
		{"primary valid", enum.ContactPrimary, true},
		{"emergency valid", enum.ContactEmergency, true},
		{"secondary valid", enum.ContactSecondary, true},
		{"work valid", enum.ContactWork, true},
		{"invalid value", enum.ContactType("unknown"), false},
		{"empty value", enum.ContactType(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseContactType(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.ContactType
		wantErr  bool
		testName string
	}{
		{"primary", enum.ContactPrimary, false, "lower primary"},
		{" Emergency ", enum.ContactEmergency, false, "trimmed emergency"},
		{"SECONDARY", enum.ContactSecondary, false, "upper secondary"},
		{"WoRk", enum.ContactWork, false, "mixed work"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseContactType(tt.in)
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

func TestContactType_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	ct := enum.ContactWork
	b, err := json.Marshal(ct)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"work\"" {
		t.Fatalf("Marshal got %s, want \"work\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.ContactType
	if err := json.Unmarshal([]byte("\" SeCoNdArY \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.ContactSecondary {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.ContactSecondary)
	}

	// Unmarshal invalid
	var u2 enum.ContactType
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid contact type, got nil")
	}
}

func TestContactType_Value(t *testing.T) {
	// Valid value
	v, err := enum.ContactPrimary.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "primary" {
		t.Fatalf("Value() got %#v, want 'primary' string", v)
	}

	// Invalid value
	var invalid enum.ContactType = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestContactType_Scan(t *testing.T) {
	// From string
	var c1 enum.ContactType
	if err := c1.Scan("EMERGENCY"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if c1 != enum.ContactEmergency {
		t.Fatalf("Scan(string) got %q, want %q", c1, enum.ContactEmergency)
	}

	// From []byte
	var c2 enum.ContactType
	if err := c2.Scan([]byte("work")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if c2 != enum.ContactWork {
		t.Fatalf("Scan([]byte) got %q, want %q", c2, enum.ContactWork)
	}

	// Invalid string value
	var c3 enum.ContactType
	if err := c3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var c4 enum.ContactType
	var src any = 3.14
	if err := c4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestContactType_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.ContactType
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
