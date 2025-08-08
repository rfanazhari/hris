package valueobject

import (
	"errors"
	"strings"
)

// EmailAddress represents an email split into username and domain parts.
// Example:
//
//	username: "username"
//	domain:   "gmail.com"
//
// Full() returns: "username@gmail.com".
//
// This value object performs light normalization and validation suitable for
// most use-cases without attempting full RFC compliance.
// - Trims spaces
// - Lowercases domain (username is kept as-is)
// - Ensures both parts are non-empty, contain no spaces, and do not contain '@'
// - Ensures domain has at least one dot and does not start/end with a dot
//
// For stricter needs, validate externally as necessary.
type EmailAddress struct {
	username string
	domain   string
}

// NewEmailAddress constructs an EmailAddress with basic normalization and validation.
func NewEmailAddress(username, domain string) (*EmailAddress, error) {
	u := strings.TrimSpace(username)
	d := strings.TrimSpace(domain)
	// domain commonly case-insensitive
	d = strings.ToLower(d)

	if u == "" {
		return nil, errors.New("username cannot be empty")
	}
	if d == "" {
		return nil, errors.New("domain cannot be empty")
	}
	if strings.ContainsAny(u, " \t\n\r@") {
		return nil, errors.New("username must not contain spaces or '@'")
	}
	if strings.ContainsAny(d, " \t\n\r@") {
		return nil, errors.New("domain must not contain spaces or '@'")
	}
	if strings.HasPrefix(d, ".") || strings.HasSuffix(d, ".") || !strings.Contains(d, ".") {
		return nil, errors.New("domain must contain a dot and not start/end with a dot")
	}

	return &EmailAddress{username: u, domain: d}, nil
}

// Username returns the local-part (before '@').
func (e EmailAddress) Username() string { return e.username }

// Domain returns the domain part (after '@').
func (e EmailAddress) Domain() string { return e.domain }

// Full returns the complete email address in the form username@domain.
func (e EmailAddress) Full() string { return e.username + "@" + e.domain }
