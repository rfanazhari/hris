package valueobject

import (
	"errors"
	"time"

	"github.com/rfanazhari/hris/domain/enum"
)

type Document struct {
	kind     enum.DocumentType
	file     FileReference
	validity ValidityPeriodDocument
}

// NewDocument constructs a Document with validation based on its fields.
// Rules:
// - kind must be a valid enum.DocumentType
// - file must be a valid FileReference (non-empty url/filename/mime)
// - validity must be a valid ValidityPeriodDocument (issued non-zero, expiry >= issued)
func NewDocument(kind enum.DocumentType, file FileReference, validity ValidityPeriodDocument) (*Document, error) {
	if !kind.Valid() {
		return nil, errors.New("invalid document type")
	}
	// Basic validation for file reference (internal fields are accessible in this package)
	if file.url == "" || file.filename == "" || file.mimeType == "" {
		return nil, errors.New("invalid file reference")
	}
	// Validate validity period constraints
	if validity.issuedDate.IsZero() {
		return nil, errors.New("issued date cannot be zero")
	}
	if validity.expiryDate != nil && validity.expiryDate.Before(validity.issuedDate) {
		return nil, errors.New("expiry date cannot be before issued date")
	}
	return &Document{kind: kind, file: file, validity: validity}, nil
}

// Kind returns the document type.
func (d Document) Kind() enum.DocumentType { return d.kind }

// File returns the associated file reference.
func (d Document) File() FileReference { return d.file }

// Validity returns the validity period value object.
func (d Document) Validity() ValidityPeriodDocument { return d.validity }

// IssuedDate convenience accessor delegates to validity.
func (d Document) IssuedDate() time.Time { return d.validity.IssuedDate() }

// ExpiryDate convenience accessor delegates to validity.
func (d Document) ExpiryDate() *time.Time { return d.validity.ExpiryDate() }

// HasExpiry reports whether the document has an expiry date.
func (d Document) HasExpiry() bool { return d.validity.HasExpiry() }

// IsExpired reports whether the document is expired at the given instant.
func (d Document) IsExpired(at time.Time) bool { return d.validity.IsExpired(at) }
