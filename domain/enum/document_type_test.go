package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestDocumentType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.DocumentType
		valid bool
	}{
		{"ktp valid", enum.DocKTP, true},
		{"npwp valid", enum.DocNPWP, true},
		{"offering_letter valid", enum.DocOfferingLetter, true},
		{"nda valid", enum.DocNDA, true},
		{"pkwt valid", enum.DocPKWT, true},
		{"other valid", enum.DocOther, true},
		{"contract_of_service valid", enum.DocContractOfService, true},
		{"scope_of_work valid", enum.DocScopeOfWork, true},
		{"tnc valid", enum.DocTnC, true},
		{"entire_agreement valid", enum.DocEntireAgreement, true},
		{"outsourcing valid", enum.DocOutsourcing, true},
		{"invalid value", enum.DocumentType("unknown"), false},
		{"empty value", enum.DocumentType(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseDocumentType(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.DocumentType
		wantErr  bool
		testName string
	}{
		{"KTP", enum.DocKTP, false, "upper KTP"},
		{" npwp ", enum.DocNPWP, false, "trimmed npwp"},
		{"OFFERING_LETTER", enum.DocOfferingLetter, false, "upper offering_letter"},
		{"NDA", enum.DocNDA, false, "upper nda"},
		{"PKWT", enum.DocPKWT, false, "upper pkwt"},
		{"Other", enum.DocOther, false, "mixed other"},
		{"Contract_Of_Service", enum.DocContractOfService, false, "contract_of_service with different case"},
		{"scope_of_work", enum.DocScopeOfWork, false, "lower scope_of_work"},
		{"TnC", enum.DocTnC, false, "mixed tnc"},
		{"ENTIRE_AGREEMENT", enum.DocEntireAgreement, false, "upper entire_agreement"},
		{"outsourcing", enum.DocOutsourcing, false, "lower outsourcing"},
		{"Contract Of Service", "", true, "invalid with spaces not underscore"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseDocumentType(tt.in)
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

func TestDocumentType_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	dt := enum.DocOfferingLetter
	b, err := json.Marshal(dt)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"offering_letter\"" {
		t.Fatalf("Marshal got %s, want \"offering_letter\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.DocumentType
	if err := json.Unmarshal([]byte("\" ENTIRE_AGREEMENT \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.DocEntireAgreement {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.DocEntireAgreement)
	}

	// Unmarshal invalid
	var u2 enum.DocumentType
	if err := json.Unmarshal([]byte("\"Contract Of Service\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid document type, got nil")
	}
}

func TestDocumentType_Value(t *testing.T) {
	// Valid value
	v, err := enum.DocKTP.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "ktp" {
		t.Fatalf("Value() got %#v, want 'ktp' string", v)
	}

	// Invalid value
	var invalid enum.DocumentType = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestDocumentType_Scan(t *testing.T) {
	// From string
	var d1 enum.DocumentType
	if err := d1.Scan("NDA"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if d1 != enum.DocNDA {
		t.Fatalf("Scan(string) got %q, want %q", d1, enum.DocNDA)
	}

	// From []byte
	var d2 enum.DocumentType
	if err := d2.Scan([]byte("scope_of_work")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if d2 != enum.DocScopeOfWork {
		t.Fatalf("Scan([]byte) got %q, want %q", d2, enum.DocScopeOfWork)
	}

	// Invalid string value
	var d3 enum.DocumentType
	if err := d3.Scan("Contract Of Service"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var d4 enum.DocumentType
	var src any = 123
	if err := d4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestDocumentType_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.DocumentType
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
