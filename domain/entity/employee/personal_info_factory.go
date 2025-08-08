package employee_entity

import (
	"errors"
	"github.com/rfanazhari/hris/domain/enum"
	vo "github.com/rfanazhari/hris/domain/valueobject"
	"time"
)

type PersonalInfoFactory struct {
	FirstName     string
	MiddleName    string
	NickName      string
	LastName      string
	BirthDate     time.Time
	PlaceOfBirth  string
	Gender        string
	Nationality   string
	MaritalStatus string
	Religion      string
}

func (f PersonalInfoFactory) Create() (*PersonalInfo, error) {
	name, err := vo.NewEmployeeName(f.FirstName, f.MiddleName, f.LastName, f.NickName)
	if err != nil {
		return nil, err
	}

	if f.BirthDate.IsZero() {
		f.BirthDate = time.Now()
	}

	if f.PlaceOfBirth == "" {
		return nil, errors.New("place of birth cannot be empty")
	}

	nationality, err := enum.ParseNationality(f.Nationality)
	if err != nil {
		return nil, err
	}

	gender, err := enum.ParseGender(f.Gender)
	if err != nil {
		return nil, err
	}

	maritalStatus, err := enum.ParseMaritalStatus(f.MaritalStatus)
	if err != nil {
		return nil, err
	}

	religion, err := enum.ParseReligion(f.Religion)
	if err != nil {
		return nil, err
	}

	return &PersonalInfo{
		name:          *name,
		birthDate:     f.BirthDate,
		placeOfBirth:  f.PlaceOfBirth,
		gender:        gender,
		nationality:   nationality,
		maritalStatus: maritalStatus,
		religion:      religion,
	}, nil
}
