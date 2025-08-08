package entity

import (
	"github.com/google/uuid"
	"github.com/rfanazhari/hris/domain/enum"
	"time"
)

type OrganizationUnit struct {
	id           uuid.UUID
	name         string
	parentUnitID *uuid.UUID
	kind         enum.OrganizationUnitKind
	createdAt    time.Time
}

func (o *OrganizationUnit) ID() uuid.UUID {
	return o.id
}

func (o *OrganizationUnit) Name() string {
	return o.name
}

func (o *OrganizationUnit) ParentID() *uuid.UUID {
	return o.parentUnitID
}

func (o *OrganizationUnit) Type() enum.OrganizationUnitKind {
	return o.kind
}

func (o *OrganizationUnit) CreatedAt() time.Time {
	return o.createdAt
}
