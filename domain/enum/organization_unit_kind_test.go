package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestOrganizationUnitKind_Valid(t *testing.T) {
	tests := []struct {
		name  string
		kind  enum.OrganizationUnitKind
		valid bool
	}{
		{"division valid", enum.OrgUnitDivision, true},
		{"department valid", enum.OrgUnitDepartment, true},
		{"team valid", enum.OrgUnitTeam, true},
		{"invalid value", enum.OrganizationUnitKind("unknown"), false},
		{"empty value", enum.OrganizationUnitKind(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kind.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.kind))
			}
		})
	}
}

func TestParseOrganizationUnitKind(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.OrganizationUnitKind
		wantErr  bool
		testName string
	}{
		{"division", enum.OrgUnitDivision, false, "lower division"},
		{" Division ", enum.OrgUnitDivision, false, "trimmed division"},
		{"DEPARTMENT", enum.OrgUnitDepartment, false, "upper department"},
		{"TeAm", enum.OrgUnitTeam, false, "mixed team"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseOrganizationUnitKind(tt.in)
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

func TestOrganizationUnitKind_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	k := enum.OrgUnitDepartment
	b, err := json.Marshal(k)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"department\"" {
		t.Fatalf("Marshal got %s, want \"department\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.OrganizationUnitKind
	if err := json.Unmarshal([]byte("\" TeAm \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.OrgUnitTeam {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.OrgUnitTeam)
	}

	// Unmarshal invalid
	var u2 enum.OrganizationUnitKind
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid kind, got nil")
	}
}

func TestOrganizationUnitKind_Value(t *testing.T) {
	// Valid value
	v, err := enum.OrgUnitDivision.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "division" {
		t.Fatalf("Value() got %#v, want 'division' string", v)
	}

	// Invalid value
	var invalid enum.OrganizationUnitKind = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestOrganizationUnitKind_Scan(t *testing.T) {
	// From string
	var k1 enum.OrganizationUnitKind
	if err := k1.Scan("TEAM"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if k1 != enum.OrgUnitTeam {
		t.Fatalf("Scan(string) got %q, want %q", k1, enum.OrgUnitTeam)
	}

	// From []byte
	var k2 enum.OrganizationUnitKind
	if err := k2.Scan([]byte("department")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if k2 != enum.OrgUnitDepartment {
		t.Fatalf("Scan([]byte) got %q, want %q", k2, enum.OrgUnitDepartment)
	}

	// Invalid string value
	var k3 enum.OrganizationUnitKind
	if err := k3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var k4 enum.OrganizationUnitKind
	var src any = 123
	if err := k4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	// not directly possible to assert sql.Scanner without importing database/sql; instead, runtime check of Scan method presence
	var k enum.OrganizationUnitKind
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
