package entity

import (
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"time"
)

// OrganizationUnit represents a unit within an organization, such as a division, department, or team.
type OrganizationUnit struct {
	id           uuid.UUID
	name         string
	parentUnitID *uuid.UUID
	kind         enum.OrganizationUnitKind
	createdAt    time.Time
}

// ID returns the unique identifier (UUID) of the OrganizationUnit.
func (o *OrganizationUnit) ID() uuid.UUID {
	return o.id
}

// Name returns the name of the organization unit.
func (o *OrganizationUnit) Name() string {
	return o.name
}

// ParentID returns the UUID of the parent organization unit or nil if it has no parent.
func (o *OrganizationUnit) ParentID() *uuid.UUID {
	return o.parentUnitID
}

// Type returns the kind of the organization unit as an instance of enum.OrganizationUnitKind.
func (o *OrganizationUnit) Type() enum.OrganizationUnitKind {
	return o.kind
}

// CreatedAt returns the timestamp indicating when the OrganizationUnit was created.
func (o *OrganizationUnit) CreatedAt() time.Time {
	return o.createdAt
}
