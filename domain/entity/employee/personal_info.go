package employee_entity

import (
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"time"
)

type PersonalInfo struct {
	name          valueobject.EmployeeName
	birthDate     time.Time
	placeOfBirth  string
	gender        enum.Gender
	nationality   enum.Nationality
	maritalStatus enum.MaritalStatus
	religion      enum.Religion
}

func (p *PersonalInfo) Name() valueobject.EmployeeName {
	return p.name
}

func (p *PersonalInfo) BirthDate() time.Time {
	return p.birthDate
}

func (p *PersonalInfo) PlaceOfBirth() string {
	return p.placeOfBirth
}

func (p *PersonalInfo) Gender() enum.Gender {
	return p.gender
}

func (p *PersonalInfo) Nationality() enum.Nationality {
	return p.nationality
}

func (p *PersonalInfo) MaritalStatus() enum.MaritalStatus {
	return p.maritalStatus
}

func (p *PersonalInfo) Religion() enum.Religion {
	return p.religion
}
