package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestRelationshipType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.RelationshipType
		valid bool
	}{
		{"wife valid", enum.RelationshipWife, true},
		{"husband valid", enum.RelationshipHusband, true},
		{"son valid", enum.RelationshipSon, true},
		{"daughter valid", enum.RelationshipDaughter, true},
		{"brother valid", enum.RelationshipBrother, true},
		{"sister valid", enum.RelationshipSister, true},
		{"father valid", enum.RelationshipFather, true},
		{"mother valid", enum.RelationshipMother, true},
		{"father_in_law valid", enum.RelationshipFatherInLaw, true},
		{"mother_in_law valid", enum.RelationshipMotherInLaw, true},
		{"grandfather valid", enum.RelationshipGrandfather, true},
		{"grandmother valid", enum.RelationshipGrandmother, true},
		{"uncle valid", enum.RelationshipUncle, true},
		{"aunt valid", enum.RelationshipAunt, true},
		{"cousin valid", enum.RelationshipCousin, true},
		{"nephew valid", enum.RelationshipNephew, true},
		{"niece valid", enum.RelationshipNiece, true},
		{"friend valid", enum.RelationshipFriend, true},
		{"partner valid", enum.RelationshipPartner, true},
		{"invalid value", enum.RelationshipType("unknown"), false},
		{"empty value", enum.RelationshipType(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseRelationshipType(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.RelationshipType
		wantErr  bool
		testName string
	}{
		{"wife", enum.RelationshipWife, false, "lower wife"},
		{" Husband ", enum.RelationshipHusband, false, "trimmed husband"},
		{"SON", enum.RelationshipSon, false, "upper son"},
		{"DaUgHtEr", enum.RelationshipDaughter, false, "mixed daughter"},
		{"BROTHER", enum.RelationshipBrother, false, "upper brother"},
		{"sister", enum.RelationshipSister, false, "lower sister"},
		{"FATHER", enum.RelationshipFather, false, "upper father"},
		{"mother", enum.RelationshipMother, false, "lower mother"},
		{" father_in_law ", enum.RelationshipFatherInLaw, false, "trimmed father_in_law"},
		{"MOTHER_IN_LAW", enum.RelationshipMotherInLaw, false, "upper mother_in_law"},
		{"grandfather", enum.RelationshipGrandfather, false, "lower grandfather"},
		{"GrandMother", enum.RelationshipGrandmother, false, "mixed grandmother"},
		{"uncle", enum.RelationshipUncle, false, "uncle"},
		{"aunt", enum.RelationshipAunt, false, "aunt"},
		{"cousin", enum.RelationshipCousin, false, "cousin"},
		{"nephew", enum.RelationshipNephew, false, "nephew"},
		{"niece", enum.RelationshipNiece, false, "niece"},
		{"friend", enum.RelationshipFriend, false, "friend"},
		{"partner", enum.RelationshipPartner, false, "partner"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseRelationshipType(tt.in)
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

func TestRelationshipType_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	r := enum.RelationshipPartner
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"partner\"" {
		t.Fatalf("Marshal got %s, want \"partner\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.RelationshipType
	if err := json.Unmarshal([]byte("\" Father_In_Law \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.RelationshipFatherInLaw {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.RelationshipFatherInLaw)
	}

	// Unmarshal invalid
	var u2 enum.RelationshipType
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid relationship type, got nil")
	}
}

func TestRelationshipType_Value(t *testing.T) {
	// Valid value
	v, err := enum.RelationshipMother.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "mother" {
		t.Fatalf("Value() got %#v, want 'mother' string", v)
	}

	// Invalid value
	var invalid enum.RelationshipType = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestRelationshipType_Scan(t *testing.T) {
	// From string
	var r1 enum.RelationshipType
	if err := r1.Scan("GRANDMOTHER"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if r1 != enum.RelationshipGrandmother {
		t.Fatalf("Scan(string) got %q, want %q", r1, enum.RelationshipGrandmother)
	}

	// From []byte
	var r2 enum.RelationshipType
	if err := r2.Scan([]byte("uncle")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if r2 != enum.RelationshipUncle {
		t.Fatalf("Scan([]byte) got %q, want %q", r2, enum.RelationshipUncle)
	}

	// Invalid string value
	var r3 enum.RelationshipType
	if err := r3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var r4 enum.RelationshipType
	var src any = 123
	if err := r4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestRelationshipType_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.RelationshipType
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
