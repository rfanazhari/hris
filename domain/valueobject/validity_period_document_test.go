package valueobject_test

import (
	"testing"
	"time"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewValidityPeriodDocument_WithoutExpiry(t *testing.T) {
	issued := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
	v, err := vo.NewValidityPeriodDocument(issued, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.IssuedDate() != issued {
		t.Fatalf("issued date mismatch")
	}
	if v.HasExpiry() {
		t.Fatalf("expected no expiry")
	}
	if v.ExpiryDate() != nil {
		t.Fatalf("expected expiry to be nil")
	}
	// Without expiry, never expired regardless of time
	if v.IsExpired(time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("expected not expired when no expiry is set")
	}
}

func TestNewValidityPeriodDocument_WithExpiryValid(t *testing.T) {
	issued := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	exp := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	v, err := vo.NewValidityPeriodDocument(issued, &exp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !v.HasExpiry() {
		t.Fatalf("expected expiry to be set")
	}
	if v.ExpiryDate() == nil || !v.ExpiryDate().Equal(exp) {
		t.Fatalf("expiry date mismatch")
	}
	// Before expiry -> not expired
	if v.IsExpired(time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)) {
		t.Fatalf("expected not expired before expiry")
	}
	// After expiry -> expired
	if !v.IsExpired(time.Date(2025, 1, 1, 0, 0, 1, 0, time.UTC)) {
		t.Fatalf("expected expired after expiry")
	}
}

func TestNewValidityPeriodDocument_InvalidZeroIssued(t *testing.T) {
	if _, err := vo.NewValidityPeriodDocument(time.Time{}, nil); err == nil {
		t.Fatalf("expected error for zero issued date")
	}
}

func TestNewValidityPeriodDocument_ExpiryBeforeIssued(t *testing.T) {
	issued := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	exp := time.Date(2024, 1, 9, 23, 59, 59, 0, time.UTC)
	if _, err := vo.NewValidityPeriodDocument(issued, &exp); err == nil {
		t.Fatalf("expected error when expiry before issued")
	}
}

func TestNewValidityPeriodDocument_ExpiryEqualIssued(t *testing.T) {
	issued := time.Date(2024, 1, 10, 10, 0, 0, 0, time.UTC)
	exp := issued
	v, err := vo.NewValidityPeriodDocument(issued, &exp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// At the exact expiry instant, not yet considered expired (strictly after)
	if v.IsExpired(exp) {
		t.Fatalf("expected not expired at the exact expiry instant")
	}
	// One nanosecond after -> expired
	if !v.IsExpired(exp.Add(time.Nanosecond)) {
		t.Fatalf("expected expired just after expiry instant")
	}
}
