package entity_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOrganizationUnitFactory_Create(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "IT Division",
			ParentUnitID: "",
			Type:         "division",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.Nil(t, err)
		assert.NotNil(t, orgUnit)
		assert.Equal(t, factory.Name, orgUnit.Name())
		assert.False(t, orgUnit.CreatedAt().IsZero())
	})
	t.Run("ValidInputWithParentID", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "IT Division",
			ParentUnitID: uuid.NewString(),
			Type:         "division",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.Nil(t, err)
		assert.NotNil(t, orgUnit)
		assert.Equal(t, factory.Name, orgUnit.Name())
		assert.False(t, orgUnit.CreatedAt().IsZero())
	})
	t.Run("InvalidID", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           "string-uuid",
			Name:         "",
			ParentUnitID: "",
			Type:         "",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, "invalid format uuid")
	})
	t.Run("InvalidParentID", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "",
			ParentUnitID: "string-uuid",
			Type:         "",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, "invalid parent unit id")
	})
	t.Run("EmptyName", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "",
			ParentUnitID: "",
			Type:         "division",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, "name cannot be empty")
	})
	t.Run("InvalidCharLengthName", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "MK",
			ParentUnitID: "",
			Type:         "division",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, "name must be at least 3 characters long")
	})
	t.Run("EmptyType", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "IT Division",
			ParentUnitID: "",
			Type:         "",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, "type cannot be empty")
	})

	t.Run("InvalidType", func(t *testing.T) {
		factory := entity.OrganizationUnitFactory{
			ID:           uuid.NewString(),
			Name:         "IT Division",
			ParentUnitID: "",
			Type:         "gudep",
			CreatedAt:    time.Time{},
		}

		orgUnit, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, orgUnit)
		assert.EqualError(t, err, fmt.Errorf("invalid OrganizationUnitKind: %q", "gudep").Error())
	})
}
