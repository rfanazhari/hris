package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestContractType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.ContractType
		valid bool
	}{
		{"pkwt valid", enum.ContractPKWT, true},
		{"pkwtt valid", enum.ContractPKWTT, true},
		{"freelance valid", enum.ContractFreelance, true},
		{"internship valid", enum.ContractInternship, true},
		{"permanent valid", enum.ContractPermanent, true},
		{"invalid value", enum.ContractType("unknown"), false},
		{"empty value", enum.ContractType(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseContractType(t *testing.T) {
	tests := []struct {
		in      string
		want    enum.ContractType
		wantErr bool
		name    string
	}{
		{"PKWT", enum.ContractPKWT, false, "upper pkwt"},
		{" pkwtt ", enum.ContractPKWTT, false, "trimmed pkwtt"},
		{"FREELANCE", enum.ContractFreelance, false, "upper freelance"},
		{"Internship", enum.ContractInternship, false, "mixed internship"},
		{"Permanent", enum.ContractPermanent, false, "mixed permanent"},
		{"Contract", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := enum.ParseContractType(tt.in)
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

func TestContractType_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	ct := enum.ContractFreelance
	b, err := json.Marshal(ct)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"freelance\"" {
		t.Fatalf("Marshal got %s, want \"freelance\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.ContractType
	if err := json.Unmarshal([]byte("\" PKWTT \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.ContractPKWTT {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.ContractPKWTT)
	}

	// Unmarshal invalid
	var u2 enum.ContractType
	if err := json.Unmarshal([]byte("\"Contract\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid contract type, got nil")
	}
}

func TestContractType_Value(t *testing.T) {
	// Valid value
	v, err := enum.ContractPKWT.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "pkwt" {
		t.Fatalf("Value() got %#v, want 'pkwt' string", v)
	}

	// Invalid value
	var invalid enum.ContractType = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestContractType_Scan(t *testing.T) {
	// From string
	var d1 enum.ContractType
	if err := d1.Scan("PKWTT"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if d1 != enum.ContractPKWTT {
		t.Fatalf("Scan(string) got %q, want %q", d1, enum.ContractPKWTT)
	}

	// From []byte
	var d2 enum.ContractType
	if err := d2.Scan([]byte("internship")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if d2 != enum.ContractInternship {
		t.Fatalf("Scan([]byte) got %q, want %q", d2, enum.ContractInternship)
	}

	// From string PERMANENT
	var d5 enum.ContractType
	if err := d5.Scan("PERMANENT"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if d5 != enum.ContractPermanent {
		t.Fatalf("Scan(string) got %q, want %q", d5, enum.ContractPermanent)
	}

	// Invalid string value
	var d3 enum.ContractType
	if err := d3.Scan("Contract"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var d4 enum.ContractType
	var src any = 123
	if err := d4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestContractType_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.ContractType
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
