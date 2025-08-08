package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type OrganizationUnitKind string

const (
	OrgUnitDivision   OrganizationUnitKind = "division"
	OrgUnitDepartment OrganizationUnitKind = "department"
	OrgUnitTeam       OrganizationUnitKind = "team"
)

func (k OrganizationUnitKind) Valid() bool {
	switch k {
	case OrgUnitDivision, OrgUnitDepartment, OrgUnitTeam:
		return true
	default:
		return false
	}
}

func ParseOrganizationUnitKind(s string) (OrganizationUnitKind, error) {
	v := OrganizationUnitKind(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid OrganizationUnitKind: %q", s)
	}
	return v, nil
}

func (k OrganizationUnitKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(k))
}

func (k *OrganizationUnitKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseOrganizationUnitKind(s)
	if err != nil {
		return err
	}
	*k = v
	return nil
}

func (k OrganizationUnitKind) Value() (driver.Value, error) {
	if !k.Valid() {
		return nil, fmt.Errorf("invalid OrganizationUnitKind: %q", k)
	}
	return string(k), nil
}

func (k *OrganizationUnitKind) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseOrganizationUnitKind(v)
		if err != nil {
			return err
		}
		*k = parsed
		return nil
	case []byte:
		return k.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for OrganizationUnitKind: %T", src)
	}
}
