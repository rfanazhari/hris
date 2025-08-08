package entity

import (
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"time"
)

// JobPosition represents a job role within an organization, storing metadata such as title, description, and salary range.
type JobPosition struct {
	id          uuid.UUID
	title       string
	description string
	gradeLevel  enum.JobGradeLevel
	salaryRange valueobject.SalaryRange
	createdAt   time.Time
}

// ID retrieves the unique identifier of the JobPosition.
func (j *JobPosition) ID() uuid.UUID {
	return j.id
}

// Title returns the title of the job position.
func (j *JobPosition) Title() string {
	return j.title
}

// Description returns the description of the job position.
func (j *JobPosition) Description() string {
	return j.description
}

// GradeLevel returns the job grade level associated with the JobPosition instance.
func (j *JobPosition) GradeLevel() enum.JobGradeLevel {
	return j.gradeLevel
}

// SalaryRange returns the salary range associated with the job position.
func (j *JobPosition) SalaryRange() valueobject.SalaryRange {
	return j.salaryRange
}

// CreatedAt returns the timestamp when the job position was created.
func (j *JobPosition) CreatedAt() time.Time {
	return j.createdAt
}
