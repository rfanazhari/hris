package valueobject_test

import (
	"net/url"
	"testing"

	vo "github.com/rfanazhari/hris/domain/valueobject"
)

func TestNewFileReference_Valid(t *testing.T) {
	f, err := vo.NewFileReference("https://storage.example.com/docs/nda-arfan.pdf", "nda-arfan.pdf", "application/pdf")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Filename() != "nda-arfan.pdf" {
		t.Fatalf("Filename mismatch: got %s", f.Filename())
	}
	if f.MimeType() != "application/pdf" {
		t.Fatalf("MimeType mismatch: got %s", f.MimeType())
	}
	// Parse the URL and check components
	u, err := url.Parse(f.URL())
	if err != nil {
		t.Fatalf("returned URL is invalid: %v", err)
	}
	if u.Scheme != "https" || u.Host != "storage.example.com" {
		t.Fatalf("unexpected URL components: scheme=%s host=%s", u.Scheme, u.Host)
	}
}

func TestNewFileReference_Normalization(t *testing.T) {
	f, err := vo.NewFileReference(" HTTPS://Storage.Example.com/docs/Doc.pdf ", "  Doc.pdf  ", "  APPLICATION/PDF  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Filename() != "Doc.pdf" {
		t.Fatalf("expected trimmed filename 'Doc.pdf', got %q", f.Filename())
	}
	if f.MimeType() != "APPLICATION/PDF" {
		t.Fatalf("expected trimmed mimeType 'APPLICATION/PDF', got %q", f.MimeType())
	}
	pu, err := url.Parse(f.URL())
	if err != nil {
		t.Fatalf("unexpected url parse error: %v", err)
	}
	if pu.Scheme != "https" {
		t.Fatalf("expected normalized scheme https, got %s", pu.Scheme)
	}
	if pu.Host != "storage.example.com" {
		t.Fatalf("expected host storage.example.com, got %s", pu.Host)
	}
}

func TestNewFileReference_InvalidInputs(t *testing.T) {
	cases := []struct {
		name     string
		url      string
		filename string
		mime     string
	}{
		{"empty url", "", "a.pdf", "application/pdf"},
		{"invalid url no scheme", "storage.example.com/a.pdf", "a.pdf", "application/pdf"},
		{"invalid url unsupported scheme", "ftp://storage.example.com/a.pdf", "a.pdf", "application/pdf"},
		{"invalid url no host", "https:///a.pdf", "a.pdf", "application/pdf"},
		{"empty filename", "https://storage.example.com/a.pdf", "", "application/pdf"},
		{"filename with slash", "https://storage.example.com/a.pdf", "dir/a.pdf", "application/pdf"},
		{"filename with backslash", "https://storage.example.com/a.pdf", "dir\\a.pdf", "application/pdf"},
		{"empty mime", "https://storage.example.com/a.pdf", "a.pdf", ""},
		{"mime without slash", "https://storage.example.com/a.pdf", "a.pdf", "applicationpdf"},
		{"mime with spaces", "https://storage.example.com/a.pdf", "a.pdf", "application/ pdf"},
		{"mime missing subtype", "https://storage.example.com/a.pdf", "a.pdf", "application/"},
		{"mime missing type", "https://storage.example.com/a.pdf", "a.pdf", "/pdf"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := vo.NewFileReference(c.url, c.filename, c.mime); err == nil {
				t.Fatalf("expected error, got nil for case %q", c.name)
			}
		})
	}
}
