package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/pkg/validation"
	"time"
)

// OrganizationUnitFactory is a factory type for creating instances of OrganizationUnit with validated properties.
type OrganizationUnitFactory struct {
	ID           string
	Name         string
	ParentUnitID string
	Type         string
	CreatedAt    time.Time
}

// Create initializes and returns a new OrganizationUnit instance or an error if validation fails.
func (f OrganizationUnitFactory) Create() (*OrganizationUnit, error) {
	var parentUnitID *uuid.UUID

	newUUID, err := uuid.Parse(f.ID)
	if err != nil {
		return nil, errors.New("invalid format uuid")
	}

	if f.ParentUnitID != "" {
		parentId, err := uuid.Parse(f.ParentUnitID)
		if err != nil {
			return nil, errors.New("invalid parent unit id")
		}
		parentUnitID = &parentId
	}

	if f.Name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(f.Name) < 3 {
		return nil, validation.CharacterLong("name", 3)
	}

	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now()
	}

	if f.Type == "" {
		return nil, errors.New("type cannot be empty")
	}

	kind, errKind := enum.ParseOrganizationUnitKind(f.Type)
	if errKind != nil {
		return nil, errKind
	}

	return &OrganizationUnit{
		id:           newUUID,
		name:         f.Name,
		parentUnitID: parentUnitID,
		kind:         kind,
		createdAt:    f.CreatedAt,
	}, nil
}
