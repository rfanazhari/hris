package valueobject

import (
	"errors"
	"net/url"
	"path/filepath"
	"strings"
)

// FileReference represents a stored file by its public URL, filename, and MIME type.
//
// Examples:
//
//	url:      "https://storage.example.com/docs/nda-arfan.pdf"
//	filename: "nda-arfan.pdf"
//	mimeType: "application/pdf"
//
// Basic normalization & validation are performed:
// - Trims spaces on all inputs
// - URL must be a valid http(s) URL with non-empty host
// - filename must be non-empty and must not contain path separators ('/' or '\\')
// - mimeType must be in the form "type/subtype" with no spaces (e.g., "image/png")
//
// For stronger rules (e.g., allowed mime types), validate externally.
type FileReference struct {
	url      string
	filename string
	mimeType string
}

// NewFileReference constructs a FileReference with basic normalization and validation.
func NewFileReference(rawURL, filename, mimeType string) (*FileReference, error) {
	u := strings.TrimSpace(rawURL)
	fn := strings.TrimSpace(filename)
	mt := strings.TrimSpace(mimeType)

	if u == "" {
		return nil, errors.New("url cannot be empty")
	}
	pu, err := url.Parse(u)
	if err != nil || pu.Scheme == "" || pu.Host == "" {
		return nil, errors.New("url must be a valid http(s) URL with host")
	}
	// Only allow http or https
	if pu.Scheme != "http" && pu.Scheme != "https" {
		return nil, errors.New("url scheme must be http or https")
	}
	// Normalize host to lowercase to ensure consistent comparisons
	pu.Host = strings.ToLower(pu.Host)

	if fn == "" {
		return nil, errors.New("filename cannot be empty")
	}
	// Reject any path separators; ensure it's a base name
	if strings.ContainsAny(fn, "/\\") {
		return nil, errors.New("filename must not contain path separators")
	}
	if filepath.Base(fn) != fn {
		return nil, errors.New("filename must be a base name")
	}

	if mt == "" {
		return nil, errors.New("mimeType cannot be empty")
	}
	if strings.ContainsAny(mt, " \t\n\r") {
		return nil, errors.New("mimeType must not contain spaces")
	}
	// very light check for type/subtype
	parts := strings.Split(mt, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, errors.New("mimeType must be in the form type/subtype")
	}

	return &FileReference{url: pu.String(), filename: fn, mimeType: mt}, nil
}

// URL returns the file absolute/public URL.
func (f FileReference) URL() string { return f.url }

// Filename returns the file name (without any path components).
func (f FileReference) Filename() string { return f.filename }

// MimeType returns the MIME type, e.g., "application/pdf".
func (f FileReference) MimeType() string { return f.mimeType }
