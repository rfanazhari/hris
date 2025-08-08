package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestContractStatus_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.ContractStatus
		valid bool
	}{
		{"active valid", enum.ContractStatusActive, true},
		{"expired valid", enum.ContractStatusExpired, true},
		{"terminated valid", enum.ContractStatusTerminated, true},
		{"invalid value", enum.ContractStatus("unknown"), false},
		{"empty value", enum.ContractStatus(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseContractStatus(t *testing.T) {
	tests := []struct {
		in      string
		want    enum.ContractStatus
		wantErr bool
		name    string
	}{
		{"ACTIVE", enum.ContractStatusActive, false, "upper active"},
		{" expired ", enum.ContractStatusExpired, false, "trimmed expired"},
		{"Terminated", enum.ContractStatusTerminated, false, "mixed terminated"},
		{"Actively", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := enum.ParseContractStatus(tt.in)
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

func TestContractStatus_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	cs := enum.ContractStatusExpired
	b, err := json.Marshal(cs)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"expired\"" {
		t.Fatalf("Marshal got %s, want \"expired\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.ContractStatus
	if err := json.Unmarshal([]byte("\" TERMINATED \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.ContractStatusTerminated {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.ContractStatusTerminated)
	}

	// Unmarshal invalid
	var u2 enum.ContractStatus
	if err := json.Unmarshal([]byte("\"Actively\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid contract status, got nil")
	}
}

func TestContractStatus_Value(t *testing.T) {
	// Valid value
	v, err := enum.ContractStatusActive.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "active" {
		t.Fatalf("Value() got %#v, want 'active' string", v)
	}

	// Invalid value
	var invalid enum.ContractStatus = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestContractStatus_Scan(t *testing.T) {
	// From string
	var c1 enum.ContractStatus
	if err := c1.Scan("EXPIRED"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if c1 != enum.ContractStatusExpired {
		t.Fatalf("Scan(string) got %q, want %q", c1, enum.ContractStatusExpired)
	}

	// From []byte
	var c2 enum.ContractStatus
	if err := c2.Scan([]byte("terminated")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if c2 != enum.ContractStatusTerminated {
		t.Fatalf("Scan([]byte) got %q, want %q", c2, enum.ContractStatusTerminated)
	}

	// Invalid string value
	var c3 enum.ContractStatus
	if err := c3.Scan("Actively"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var c4 enum.ContractStatus
	var src any = 123
	if err := c4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestContractStatus_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.ContractStatus
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
