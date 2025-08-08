package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestEmploymentStatus_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.EmploymentStatus
		valid bool
	}{
		{"active valid", enum.EmploymentActive, true},
		{"resigned valid", enum.EmploymentResigned, true},
		{"on_leave valid", enum.EmploymentOnLeave, true},
		{"invalid value", enum.EmploymentStatus("unknown"), false},
		{"empty value", enum.EmploymentStatus(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseEmploymentStatus(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.EmploymentStatus
		wantErr  bool
		testName string
	}{
		{"ACTIVE", enum.EmploymentActive, false, "upper active"},
		{" resigned ", enum.EmploymentResigned, false, "trimmed resigned"},
		{"On_Leave", enum.EmploymentOnLeave, false, "mixed on_leave with underscore"},
		{"On Leave", "", true, "invalid with space instead of underscore"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseEmploymentStatus(tt.in)
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

func TestEmploymentStatus_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	es := enum.EmploymentOnLeave
	b, err := json.Marshal(es)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"on_leave\"" {
		t.Fatalf("Marshal got %s, want \"on_leave\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.EmploymentStatus
	if err := json.Unmarshal([]byte("\" ACTIVE \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.EmploymentActive {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.EmploymentActive)
	}

	// Unmarshal invalid (space instead of underscore)
	var u2 enum.EmploymentStatus
	if err := json.Unmarshal([]byte("\"On Leave\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid employment status, got nil")
	}
}

func TestEmploymentStatus_Value(t *testing.T) {
	// Valid value
	v, err := enum.EmploymentResigned.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "resigned" {
		t.Fatalf("Value() got %#v, want 'resigned' string", v)
	}

	// Invalid value
	var invalid enum.EmploymentStatus = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestEmploymentStatus_Scan(t *testing.T) {
	// From string
	var e1 enum.EmploymentStatus
	if err := e1.Scan("ON_LEAVE"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if e1 != enum.EmploymentOnLeave {
		t.Fatalf("Scan(string) got %q, want %q", e1, enum.EmploymentOnLeave)
	}

	// From []byte
	var e2 enum.EmploymentStatus
	if err := e2.Scan([]byte("active")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if e2 != enum.EmploymentActive {
		t.Fatalf("Scan([]byte) got %q, want %q", e2, enum.EmploymentActive)
	}

	// Invalid string value
	var e3 enum.EmploymentStatus
	if err := e3.Scan("On Leave"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var e4 enum.EmploymentStatus
	var src any = 123
	if err := e4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestEmploymentStatus_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.EmploymentStatus
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
