package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type JobGradeLevel string

const (
	JobGradeIntern   JobGradeLevel = "intern"
	JobGradeJunior   JobGradeLevel = "junior"
	JobGradeMid      JobGradeLevel = "mid"
	JobGradeSenior   JobGradeLevel = "senior"
	JobGradeLead     JobGradeLevel = "lead"
	JobGradeManager  JobGradeLevel = "manager"
	JobGradeDirector JobGradeLevel = "director"
)

func (j JobGradeLevel) Valid() bool {
	switch j {
	case JobGradeIntern, JobGradeJunior, JobGradeMid, JobGradeSenior, JobGradeLead, JobGradeManager, JobGradeDirector:
		return true
	default:
		return false
	}
}

func ParseJobGradeLevel(s string) (JobGradeLevel, error) {
	v := JobGradeLevel(strings.ToLower(strings.TrimSpace(s)))
	if !v.Valid() {
		return "", fmt.Errorf("invalid JobGradeLevel: %q", s)
	}
	return v, nil
}

func (j JobGradeLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(j))
}

func (j *JobGradeLevel) UnmarshalJSON(b []byte) error {
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

func (j JobGradeLevel) Value() (driver.Value, error) {
	if !j.Valid() {
		return nil, fmt.Errorf("invalid JobGradeLevel: %q", j)
	}
	return string(j), nil
}

func (j *JobGradeLevel) Scan(src any) error {
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
		return fmt.Errorf("unsupported scan type for JobGradeLevel: %T", src)
	}
}
