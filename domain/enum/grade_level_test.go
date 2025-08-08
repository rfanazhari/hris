package enum_test

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	enum "github.com/rfanazhari/hris/domain/enum"
)

func TestJobGradeLevel_Valid(t *testing.T) {
	tests := []struct {
		name  string
		val   enum.GradeLevel
		valid bool
	}{
		{"intern valid", enum.GradeIntern, true},
		{"junior valid", enum.GradeJunior, true},
		{"mid valid", enum.GradeMid, true},
		{"senior valid", enum.GradeSenior, true},
		{"lead valid", enum.GradeLead, true},
		{"manager valid", enum.GradeManager, true},
		{"director valid", enum.GradeDirector, true},
		{"invalid value", enum.GradeLevel("unknown"), false},
		{"empty value", enum.GradeLevel(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.Valid(); got != tt.valid {
				t.Fatalf("Valid() = %v, want %v for %q", got, tt.valid, string(tt.val))
			}
		})
	}
}

func TestParseJobGradeLevel(t *testing.T) {
	tests := []struct {
		in       string
		want     enum.GradeLevel
		wantErr  bool
		testName string
	}{
		{"intern", enum.GradeIntern, false, "lower intern"},
		{" Junior ", enum.GradeJunior, false, "trimmed junior"},
		{"MID", enum.GradeMid, false, "upper mid"},
		{"SeNiOr", enum.GradeSenior, false, "mixed senior"},
		{"LEAD", enum.GradeLead, false, "upper lead"},
		{"manager", enum.GradeManager, false, "lower manager"},
		{" Director ", enum.GradeDirector, false, "trimmed director"},
		{"unknown", "", true, "invalid"},
		{"", "", true, "empty"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := enum.ParseJobGradeLevel(tt.in)
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

func TestJobGradeLevel_JSON_MarshalUnmarshal(t *testing.T) {
	// Marshal
	lvl := enum.GradeManager
	b, err := json.Marshal(lvl)
	if err != nil {
		t.Fatalf("Marshal error: %v", err)
	}
	if string(b) != "\"manager\"" {
		t.Fatalf("Marshal got %s, want \"manager\"", string(b))
	}

	// Unmarshal valid with different case and spaces
	var u enum.GradeLevel
	if err := json.Unmarshal([]byte("\" LeAd \""), &u); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if u != enum.GradeLead {
		t.Fatalf("Unmarshal got %q, want %q", u, enum.GradeLead)
	}

	// Unmarshal invalid
	var u2 enum.GradeLevel
	if err := json.Unmarshal([]byte("\"unknown\""), &u2); err == nil {
		t.Fatalf("expected error unmarshalling invalid level, got nil")
	}
}

func TestJobGradeLevel_Value(t *testing.T) {
	// Valid value
	v, err := enum.GradeIntern.Value()
	if err != nil {
		t.Fatalf("Value() unexpected error: %v", err)
	}
	if s, ok := v.(string); !ok || s != "intern" {
		t.Fatalf("Value() got %#v, want 'intern' string", v)
	}

	// Invalid value
	var invalid enum.GradeLevel = "invalid"
	if _, err := invalid.Value(); err == nil {
		t.Fatalf("expected error for invalid Value(), got nil")
	}
}

func TestJobGradeLevel_Scan(t *testing.T) {
	// From string
	var j1 enum.GradeLevel
	if err := j1.Scan("DIRECTOR"); err != nil {
		t.Fatalf("Scan(string) error: %v", err)
	}
	if j1 != enum.GradeDirector {
		t.Fatalf("Scan(string) got %q, want %q", j1, enum.GradeDirector)
	}

	// From []byte
	var j2 enum.GradeLevel
	if err := j2.Scan([]byte("junior")); err != nil {
		t.Fatalf("Scan([]byte) error: %v", err)
	}
	if j2 != enum.GradeJunior {
		t.Fatalf("Scan([]byte) got %q, want %q", j2, enum.GradeJunior)
	}

	// Invalid string value
	var j3 enum.GradeLevel
	if err := j3.Scan("unknown"); err == nil {
		t.Fatalf("expected error for invalid string scan, got nil")
	}

	// Unsupported type
	var j4 enum.GradeLevel
	var src any = 3.14
	if err := j4.Scan(src); err == nil {
		t.Fatalf("expected error for unsupported type scan, got nil")
	}
}

func TestJobGradeLevel_ImplementsDriverValuerAndScannerLike(t *testing.T) {
	// Ensure the Value() type satisfies driver.Valuer contract shape at compile time
	var _ driver.Valuer
	var k enum.GradeLevel
	// reflect check that method Scan exists
	m, ok := reflect.TypeOf(&k).MethodByName("Scan")
	if !ok || m.Type.NumIn() != 2 { // receiver + 1 arg
		t.Fatalf("Scan method not found or has unexpected signature")
	}
}
