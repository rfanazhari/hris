package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"github.com/rfanazhari/hris/pkg/validation"
	"time"
)

// JobPositionFactory represents a factory for creating JobPosition objects with validation rules applied.
type JobPositionFactory struct {
	ID             string
	Title          string
	Description    string
	GradeLevel     string
	SalaryMin      int64
	SalaryMax      int64
	SalaryCurrency string
	CreatedAt      time.Time
}

// Create generates a new JobPosition using the factory data, validating fields like ID, Title, Description, and Salary.
func (f JobPositionFactory) Create() (*JobPosition, error) {
	newUUID, err := uuid.Parse(f.ID)
	if err != nil {
		return nil, errors.New("invalid format uuid")
	}
	if f.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if len(f.Title) < 3 {
		return nil, validation.CharacterLong("title", 3)
	}

	if f.Description == "" {
		return nil, errors.New("description cannot be empty")
	}
	if len(f.Description) < 50 {
		return nil, validation.CharacterLong("description", 50)
	}

	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now()
	}

	grade, errGrade := enum.ParseJobGradeLevel(f.GradeLevel)
	if errGrade != nil {
		return nil, errGrade
	}
	salaryRange, errSalary := valueobject.NewSalaryRange(f.SalaryMin, f.SalaryMax, f.SalaryCurrency)
	if errSalary != nil {
		return nil, errSalary
	}

	return &JobPosition{
		id:          newUUID,
		title:       f.Title,
		description: f.Description,
		gradeLevel:  grade,
		salaryRange: *salaryRange,
		createdAt:   f.CreatedAt,
	}, nil
}
