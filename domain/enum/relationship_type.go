package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// RelationshipType represents familial or personal relationship types for dependents/contacts.
// Implements json (un)marshaling and database/sql interfaces.
// Use ParseRelationshipType to safely convert from string.
type RelationshipType string

const (
	RelationshipWife        RelationshipType = "wife"
	RelationshipHusband     RelationshipType = "husband"
	RelationshipSon         RelationshipType = "son"
	RelationshipDaughter    RelationshipType = "daughter"
	RelationshipBrother     RelationshipType = "brother"
	RelationshipSister      RelationshipType = "sister"
	RelationshipFather      RelationshipType = "father"
	RelationshipMother      RelationshipType = "mother"
	RelationshipFatherInLaw RelationshipType = "father_in_law"
	RelationshipMotherInLaw RelationshipType = "mother_in_law"
	RelationshipGrandfather RelationshipType = "grandfather"
	RelationshipGrandmother RelationshipType = "grandmother"
	RelationshipUncle       RelationshipType = "uncle"
	RelationshipAunt        RelationshipType = "aunt"
	RelationshipCousin      RelationshipType = "cousin"
	RelationshipNephew      RelationshipType = "nephew"
	RelationshipNiece       RelationshipType = "niece"
	RelationshipFriend      RelationshipType = "friend"
	RelationshipPartner     RelationshipType = "partner"
)

func (r RelationshipType) Valid() bool {
	switch r {
	case RelationshipWife,
		RelationshipHusband,
		RelationshipSon,
		RelationshipDaughter,
		RelationshipBrother,
		RelationshipSister,
		RelationshipFather,
		RelationshipMother,
		RelationshipFatherInLaw,
		RelationshipMotherInLaw,
		RelationshipGrandfather,
		RelationshipGrandmother,
		RelationshipUncle,
		RelationshipAunt,
		RelationshipCousin,
		RelationshipNephew,
		RelationshipNiece,
		RelationshipFriend,
		RelationshipPartner:
		return true
	default:
		return false
	}
}

func ParseRelationshipType(s string) (RelationshipType, error) {
	v := RelationshipType(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid RelationshipType: %q", s)
	}
	return v, nil
}

func (r RelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *RelationshipType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseRelationshipType(s)
	if err != nil {
		return err
	}
	*r = v
	return nil
}

func (r RelationshipType) Value() (driver.Value, error) {
	if !r.Valid() {
		return nil, fmt.Errorf("invalid RelationshipType: %q", r)
	}
	return string(r), nil
}

func (r *RelationshipType) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseRelationshipType(v)
		if err != nil {
			return err
		}
		*r = parsed
		return nil
	case []byte:
		return r.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for RelationshipType: %T", src)
	}
}
