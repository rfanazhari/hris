package valueobject_test

import (
	"encoding/json"
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewSalaryRange_ValidCurrencies(t *testing.T) {
	tests := []struct {
		name     string
		min      int64
		max      int64
		currency string
	}{
		{"IDR exact", 0, 10000000, "IDR"},
		{"USD lowercase normalized", 50000, 150000, "usd"},
		{"IDR spaced", 1, 2, " idr "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := vo.NewSalaryRange(tt.min, tt.max, tt.currency)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if s.Min != tt.min || s.Max != tt.max {
				t.Fatalf("unexpected min/max: got %d-%d", s.Min, s.Max)
			}
			if s.Currency != "IDR" && s.Currency != "USD" {
				t.Fatalf("unexpected currency normalization: %s", s.Currency)
			}
		})
	}
}

func TestNewSalaryRange_InvalidValues(t *testing.T) {
	cases := []struct {
		name string
		min  int64
		max  int64
		curr string
	}{
		{"negative min", -1, 0, "IDR"},
		{"max less than min", 10, 5, "USD"},
		{"empty currency", 0, 0, ""},
		{"unsupported currency", 0, 0, "EUR"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewSalaryRange(c.min, c.max, c.curr); err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}

func TestSalaryRange_JSON(t *testing.T) {
	s := vo.SalaryRange{Min: 1000, Max: 2000, Currency: "usd"}
	// Marshal should normalize Currency
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal back to map error: %v", err)
	}
	if out["currency"].(string) != "USD" {
		t.Fatalf("currency not normalized on marshal, got %v", out["currency"])
	}

	// Unmarshal with lenient input and validate
	var s2 vo.SalaryRange
	if err := json.Unmarshal([]byte(`{"min": 0, "max": 10, "currency": " idr "}`), &s2); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if s2.Currency != "IDR" {
		t.Fatalf("unmarshal normalization failed, got %s", s2.Currency)
	}

	// Unmarshal invalid (max < min)
	var s3 vo.SalaryRange
	if err := json.Unmarshal([]byte(`{"min": 10, "max": 1, "currency": "USD"}`), &s3); err == nil {
		t.Fatalf("expected error for invalid range, got nil")
	}
}
