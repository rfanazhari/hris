package valueobject

import (
	"errors"
	"time"
)

// ValidityPeriodDocument represents the validity period of a document.
//
// Fields:
//   - issuedDate: the date/time when the document was issued (must be non-zero)
//   - expiryDate: optional expiration date/time (nil means no expiration)
//
// Rules:
//   - expiryDate, if provided, must be the same as or after issuedDate.
//   - IsExpired(at) returns true only if there is an expiryDate and the
//     provided time instant is strictly after it.
//
// Notes:
//   - Time zone handling is delegated to callers; this type stores the times
//     as provided.
//   - For current-time checks, pass time.Now() explicitly to IsExpired.
type ValidityPeriodDocument struct {
	issuedDate time.Time
	expiryDate *time.Time
}

// NewValidityPeriodDocument constructs a ValidityPeriodDocument with validation.
func NewValidityPeriodDocument(issuedDate time.Time, expiryDate *time.Time) (*ValidityPeriodDocument, error) {
	if issuedDate.IsZero() {
		return nil, errors.New("issued date cannot be zero")
	}
	if expiryDate != nil {
		// ensure expiry is not before issued
		if expiryDate.Before(issuedDate) {
			return nil, errors.New("expiry date cannot be before issued date")
		}
	}
	return &ValidityPeriodDocument{issuedDate: issuedDate, expiryDate: expiryDate}, nil
}

// IssuedDate returns the issue time of the document.
func (v ValidityPeriodDocument) IssuedDate() time.Time { return v.issuedDate }

// ExpiryDate returns the pointer to expiry time if any; nil means no expiration.
func (v ValidityPeriodDocument) ExpiryDate() *time.Time { return v.expiryDate }

// HasExpiry reports whether the document has an expiry date.
func (v ValidityPeriodDocument) HasExpiry() bool { return v.expiryDate != nil }

// IsExpired reports whether the document is expired at the given instant.
// If there is no expiry date, it always returns false.
// Expiration is considered to happen strictly after the expiry instant.
func (v ValidityPeriodDocument) IsExpired(at time.Time) bool {
	if v.expiryDate == nil {
		return false
	}
	return at.After(*v.expiryDate)
}
