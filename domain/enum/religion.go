package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// Religion represents a person's religion/belief.
// Allowed values (string representation):
// - "Islam"
// - "Kristen Protestan"
// - "Katolik"
// - "Hindu"
// - "Buddha"
// - "Konghucu"
// - "Lainnya"
// - "Tidak Ada"
// Use ParseReligion to safely convert from string (case-insensitive, trims spaces).
// Implements json (un)marshaling and database/sql interfaces.
type Religion string

const (
	ReligionIslam      Religion = "islam"
	ReligionProtestant Religion = "kristen protestan"
	ReligionCatholic   Religion = "katolik"
	ReligionHindu      Religion = "hindu"
	ReligionBuddha     Religion = "buddha"
	ReligionKonghucu   Religion = "konghucu"
	ReligionOther      Religion = "lainnya"
	ReligionNone       Religion = "tidak ada"
)

func (r Religion) Valid() bool {
	switch r {
	case ReligionIslam,
		ReligionProtestant,
		ReligionCatholic,
		ReligionHindu,
		ReligionBuddha,
		ReligionKonghucu,
		ReligionOther,
		ReligionNone:
		return true
	default:
		return false
	}
}

func ParseReligion(s string) (Religion, error) {
	in := strings.TrimSpace(s)
	switch {
	case strings.EqualFold(in, string(ReligionIslam)):
		return ReligionIslam, nil
	case strings.EqualFold(in, string(ReligionProtestant)):
		return ReligionProtestant, nil
	case strings.EqualFold(in, string(ReligionCatholic)):
		return ReligionCatholic, nil
	case strings.EqualFold(in, string(ReligionHindu)):
		return ReligionHindu, nil
	case strings.EqualFold(in, string(ReligionBuddha)):
		return ReligionBuddha, nil
	case strings.EqualFold(in, string(ReligionKonghucu)):
		return ReligionKonghucu, nil
	case strings.EqualFold(in, string(ReligionOther)):
		return ReligionOther, nil
	case strings.EqualFold(in, string(ReligionNone)):
		return ReligionNone, nil
	default:
		return "", fmt.Errorf("invalid Religion: %q", s)
	}
}

func (r Religion) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *Religion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseReligion(s)
	if err != nil {
		return err
	}
	*r = v
	return nil
}

func (r Religion) Value() (driver.Value, error) {
	if !r.Valid() {
		return nil, fmt.Errorf("invalid Religion: %q", r)
	}
	return string(r), nil
}

func (r *Religion) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseReligion(v)
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case []byte:
		return r.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for Religion: %T", src)
	}
}
