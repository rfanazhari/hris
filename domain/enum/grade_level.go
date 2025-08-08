package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type GradeLevel string

const (
	GradeIntern   GradeLevel = "intern"
	GradeJunior   GradeLevel = "junior"
	GradeMid      GradeLevel = "mid"
	GradeSenior   GradeLevel = "senior"
	GradeLead     GradeLevel = "lead"
	GradeManager  GradeLevel = "manager"
	GradeDirector GradeLevel = "director"
)

func (j GradeLevel) Valid() bool {
	switch j {
	case GradeIntern, GradeJunior, GradeMid, GradeSenior, GradeLead, GradeManager, GradeDirector:
		return true
	default:
		return false
	}
}

func ParseJobGradeLevel(s string) (GradeLevel, error) {
	v := GradeLevel(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid GradeLevel: %q", s)
	}
	return v, nil
}

func (j GradeLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(j))
}

func (j *GradeLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := ParseJobGradeLevel(s)
	if err != nil {
		return err
	}
	*j = v
	return nil
}

func (j GradeLevel) Value() (driver.Value, error) {
	if !j.Valid() {
		return nil, fmt.Errorf("invalid GradeLevel: %q", j)
	}
	return string(j), nil
}

func (j *GradeLevel) Scan(src any) error {
	switch v := src.(type) {
	case string:
		parsed, err := ParseJobGradeLevel(v)
		if err != nil {
			return err
		}
		*j = parsed
		return nil
	case []byte:
		return j.Scan(string(v))
	default:
		return fmt.Errorf("unsupported scan type for GradeLevel: %T", src)
	}
}
