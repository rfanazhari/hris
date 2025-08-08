package entity

import (
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"github.com/rfanazhari/hris/pkg/fake"
	"github.com/rfanazhari/hris/pkg/validation"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJobPositionFactory_Create(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "Developer",
			Description:    fake.Paragraph(1, 12),
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		grade, _ := enum.ParseJobGradeLevel(factory.GradeLevel)
		assert.Nil(t, err)
		assert.NotNil(t, jobPosition)
		assert.Equal(t, factory.Title, jobPosition.Title())
		assert.Equal(t, factory.Description, jobPosition.Description())
		assert.Equal(t, grade, jobPosition.GradeLevel())
		assert.False(t, jobPosition.CreatedAt().IsZero())
	})
	t.Run("InvalidID", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             "uuid",
			Title:          "",
			Description:    "",
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, "invalid format uuid")
	})
	t.Run("EmptyTitle", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "",
			Description:    "",
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, "title cannot be empty")
	})
	t.Run("InvalidCharLengthTitle", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "di",
			Description:    "",
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, validation.CharacterLong("title", 3).Error())
	})
	t.Run("EmptyDescription", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "Developer",
			Description:    "",
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, "description cannot be empty")
	})

	t.Run("InvalidCharLengthDescription", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "Developer",
			Description:    "I am a developer",
			GradeLevel:     "junior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, validation.CharacterLong("description", 50).Error())
	})

	t.Run("InvalidGrade", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "Developer",
			Description:    fake.Paragraph(1, 12),
			GradeLevel:     "yunior",
			SalaryMin:      1000,
			SalaryMax:      10000,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		_, errGrade := enum.ParseJobGradeLevel(factory.GradeLevel)

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, errGrade.Error())
	})

	t.Run("InvalidSalaryRange", func(t *testing.T) {
		factory := JobPositionFactory{
			ID:             uuid.NewString(),
			Title:          "Developer",
			Description:    fake.Paragraph(1, 12),
			GradeLevel:     "junior",
			SalaryMin:      -1,
			SalaryMax:      0,
			SalaryCurrency: "idr",
			CreatedAt:      time.Now(),
		}

		_, errSalary := valueobject.NewSalaryRange(factory.SalaryMin, factory.SalaryMax, factory.SalaryCurrency)

		jobPosition, err := factory.Create()
		assert.NotNil(t, err)
		assert.Nil(t, jobPosition)
		assert.EqualError(t, err, errSalary.Error())
	})
}
