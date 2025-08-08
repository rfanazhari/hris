package valueobject_test

import (
	"testing"
	"time"

	enum "github.com/rfanazhari/hris/domain/enum"
	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewDocument_Valid(t *testing.T) {
	f, err := vo.NewFileReference("https://storage.example.com/docs/contract.pdf", "contract.pdf", "application/pdf")
	if err != nil {
		t.Fatalf("unexpected file ref error: %v", err)
	}
	issued := time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC)
	exp := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	vd, err := vo.NewValidityPeriodDocument(issued, &exp)
	if err != nil {
		t.Fatalf("unexpected validity error: %v", err)
	}
	d, err := vo.NewDocument(enum.DocContractOfService, *f, *vd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.Kind() != enum.DocContractOfService {
		t.Fatalf("kind mismatch")
	}
	if d.File().URL() != f.URL() || d.File().Filename() != f.Filename() || d.File().MimeType() != f.MimeType() {
		t.Fatalf("file reference mismatch")
	}
	if !d.Validity().HasExpiry() {
		t.Fatalf("expected validity to have expiry")
	}
	if d.IssuedDate() != issued {
		t.Fatalf("issued date mismatch")
	}
	if d.ExpiryDate() == nil || !d.ExpiryDate().Equal(exp) {
		t.Fatalf("expiry date mismatch")
	}
	// Delegation checks
	if d.IsExpired(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("should not be expired yet")
	}
	if !d.IsExpired(exp.Add(time.Nanosecond)) {
		t.Fatalf("should be expired after expiry instant")
	}
}

func TestNewDocument_InvalidType(t *testing.T) {
	f, _ := vo.NewFileReference("https://example.com/a.pdf", "a.pdf", "application/pdf")
	issued := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	vd, _ := vo.NewValidityPeriodDocument(issued, nil)
	if _, err := vo.NewDocument(enum.DocumentType("invalid"), *f, *vd); err == nil {
		t.Fatalf("expected error for invalid document type")
	}
}

func TestNewDocument_InvalidFile(t *testing.T) {
	// zero value FileReference is invalid
	var zeroFile vo.FileReference
	issued := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	vd, _ := vo.NewValidityPeriodDocument(issued, nil)
	if _, err := vo.NewDocument(enum.DocNDA, zeroFile, *vd); err == nil {
		t.Fatalf("expected error for invalid file reference")
	}
}

func TestNewDocument_InvalidValidity(t *testing.T) {
	f, _ := vo.NewFileReference("https://example.com/a.pdf", "a.pdf", "application/pdf")
	// invalid: zero issued date
	var invalidVD vo.ValidityPeriodDocument // zero value has zero issued date
	if _, err := vo.NewDocument(enum.DocNDA, *f, invalidVD); err == nil {
		t.Fatalf("expected error for invalid validity (zero issued)")
	}
}
