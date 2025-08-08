package entity

import (
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"time"
)

type JobPosition struct {
	id          uuid.UUID
	title       string
	description string
	gradeLevel  enum.JobGradeLevel
	salaryRange valueobject.SalaryRange
	createdAt   time.Time
}

func (j *JobPosition) ID() uuid.UUID {
	return j.id
}

func (j *JobPosition) Title() string {
	return j.title
}

func (j *JobPosition) Description() string {
	return j.description
}

func (j *JobPosition) GradeLevel() enum.JobGradeLevel {
	return j.gradeLevel
}

func (j *JobPosition) SalaryRange() valueobject.SalaryRange {
	return j.salaryRange
}

func (j *JobPosition) CreatedAt() time.Time {
	return j.createdAt
}
