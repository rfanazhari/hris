package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestReligion_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.Religion
		valid bool
	}{
		{"Islam valid", enum.ReligionIslam, true},
		{"Protestant valid", enum.ReligionProtestant, true},
		{"Catholic valid", enum.ReligionCatholic, true},
		{"Hindu valid", enum.ReligionHindu, true},
		{"Buddha valid", enum.ReligionBuddha, true},
		{"Konghucu valid", enum.ReligionKonghucu, true},
		{"Other valid", enum.ReligionOther, true},
		{"None valid", enum.ReligionNone, true},
		{"invalid value", enum.Religion("unknown"), false},
		{"empty value", enum.Religion(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseReligion(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.Religion
		wantErr  bool
		testName string
	}{
		{"islam", enum.ReligionIslam, false, "lower islam"},
		{" Kristen Protestan ", enum.ReligionProtestant, false, "trimmed protestant"},
		{"KATOLIK", enum.ReligionCatholic, false, "upper catholic"},
		{"hInDu", enum.ReligionHindu, false, "mixed hindu"},
		{"buddha", enum.ReligionBuddha, false, "lower buddha"},
		{"KONGHUCU", enum.ReligionKonghucu, false, "upper konghucu"},
		{"LAINNYA", enum.ReligionOther, false, "upper lainnya"},
		{" tidak ada ", enum.ReligionNone, false, "trimmed none"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseReligion(tt.in)
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

func TestReligion_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	r := enum.ReligionProtestant
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"Kristen Protestan\"" {
		t.Fatalf("Marshal got %s, want \"Kristen Protestan\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.Religion
	if err := json.Unmarshal([]byte("\" lAiNnYa \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.ReligionOther {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.ReligionOther)
	}

	// Unmarshal invalid
	var u2 enum.Religion
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid religion, got nil")
	}
}

func TestReligion_Value(t *testing.T) {
	// Valid value
	v, err := enum.ReligionIslam.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "Islam" {
		t.Fatalf("Value() got %#v, want 'Islam' string", v)
	}

	// Invalid value
	var invalid enum.Religion = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestReligion_Scan(t *testing.T) {
	// From string
	var r1 enum.Religion
	if err := r1.Scan("katolik"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if r1 != enum.ReligionCatholic {
		t.Fatalf("Scan(string) got %q, want %q", r1, enum.ReligionCatholic)
	}

	// From []byte
	var r2 enum.Religion
	if err := r2.Scan([]byte("KONGHUCU")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if r2 != enum.ReligionKonghucu {
		t.Fatalf("Scan([]byte) got %q, want %q", r2, enum.ReligionKonghucu)
	}

	// Invalid string value
	var r3 enum.Religion
	if err := r3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var r4 enum.Religion
	var src any = 3.14
	if err := r4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestReligion_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.Religion
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
