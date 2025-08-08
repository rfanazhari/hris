package valueobject

import (
	"errors"
	"strings"
)

// EmployeeName represents a person's name broken into common parts.
// All parts are stored trimmed. Only firstName and lastName are required.
//
// Fields:
//   - firstName (required)
//   - middleName (optional)
//   - lastName (required)
//   - nickName (optional)
//
// FullName() returns a space-joined combination of firstName, middleName (if any),
// and lastName, without extra spaces.
// Example:
//   firstName: "John", middleName: "Ronald", lastName: "Reuel" -> "John Ronald Reuel"
//   firstName: "Jane", middleName: "", lastName: "Doe" -> "Jane Doe"
//
// NickName is provided as an extra label but is not included in FullName.
// Use NickName() or choose externally how to represent preferred names.
//
// This value object keeps light validation; add stronger domain rules externally if needed.
type EmployeeName struct {
	firstName  string
	middleName string
	lastName   string
	nickName   string
}

// NewEmployeeName constructs an EmployeeName with basic normalization and validation.
// - Trims spaces on all fields
// - Requires firstName and lastName to be non-empty
func NewEmployeeName(firstName, middleName, lastName, nickName string) (*EmployeeName, error) {
	fn := strings.TrimSpace(firstName)
	mn := strings.TrimSpace(middleName)
	ln := strings.TrimSpace(lastName)
	nn := strings.TrimSpace(nickName)

	if fn == "" {
		return nil, errors.New("first name cannot be empty")
	}
	if ln == "" {
		return nil, errors.New("last name cannot be empty")
	}

	return &EmployeeName{
		firstName:  fn,
		middleName: mn,
		lastName:   ln,
		nickName:   nn,
	}, nil
}

// FirstName returns the first name.
func (e EmployeeName) FirstName() string { return e.firstName }

// MiddleName returns the middle name (may be empty).
func (e EmployeeName) MiddleName() string { return e.middleName }

// LastName returns the last name.
func (e EmployeeName) LastName() string { return e.lastName }

// NickName returns the nickname (may be empty).
func (e EmployeeName) NickName() string { return e.nickName }

// FullName returns the combination of firstName, middleName (if present) and lastName,
// joined by single spaces without leading/trailing spaces.
func (e EmployeeName) FullName() string {
	parts := make([]string, 0, 3)
	if e.firstName != "" {
		parts = append(parts, e.firstName)
	}
	if e.middleName != "" {
		parts = append(parts, e.middleName)
	}
	if e.lastName != "" {
		parts = append(parts, e.lastName)
	}
	return strings.Join(parts, " ")
}
