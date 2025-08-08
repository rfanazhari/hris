package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// DocumentType represents various types of documents used in HR processes.
// Allowed values (string representation):
// - "ktp"               // Kartu Tanda Penduduk
// - "npwp"              // Nomor Pokok Wajib Pajak
// - "offering_letter"   // Surat Penawaran Kerja
// - "nda"               // Non-Disclosure Agreement
// - "pkwt"              // Perjanjian Kerja Waktu Tertentu
// - "other"             // Untuk dokumen lain yang belum didefinisikan
// - "contract_of_service"
// - "scope_of_work"
// - "tnc"               // term and condition
// - "entire_agreement"
// - "outsourcing"
//
// Use ParseDocumentType to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type DocumentType string

const (
	DocKTP               DocumentType = "ktp"
	DocNPWP              DocumentType = "npwp"
	DocOfferingLetter    DocumentType = "offering_letter"
	DocNDA               DocumentType = "nda"
	DocPKWT              DocumentType = "pkwt"
	DocOther             DocumentType = "other"
	DocContractOfService DocumentType = "contract_of_service"
	DocScopeOfWork       DocumentType = "scope_of_work"
	DocTnC               DocumentType = "tnc"
	DocEntireAgreement   DocumentType = "entire_agreement"
	DocOutsourcing       DocumentType = "outsourcing"
)

func (d DocumentType) Valid() bool {
	switch d {
	case DocKTP, DocNPWP, DocOfferingLetter, DocNDA, DocPKWT, DocOther,
		DocContractOfService, DocScopeOfWork, DocTnC, DocEntireAgreement, DocOutsourcing:
		return true
	default:
		return false
	}
}

func ParseDocumentType(s string) (DocumentType, error) {
	v := DocumentType(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid DocumentType: %q", s)
	}
	return v, nil
}

func (d DocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

func (d *DocumentType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseDocumentType(s)
	if err != nil {
		return err
	}
	*d = v
	return nil
}

func (d DocumentType) Value() (driver.Value, error) {
	if !d.Valid() {
		return nil, fmt.Errorf("invalid DocumentType: %q", d)
	}
	return string(d), nil
}

func (d *DocumentType) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseDocumentType(v)
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	case []byte:
		return d.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for DocumentType: %T", src)
	}
}
