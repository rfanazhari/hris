package employee_entity_test

import (
	"github.com/go-faker/faker/v4"
	employee_entity "github.com/rfanazhari/hris/domain/entity/employee"
	"github.com/rfanazhari/hris/domain/enum"
	"github.com/rfanazhari/hris/domain/valueobject"
	"github.com/rfanazhari/hris/pkg/fake"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPersonalInfoFactory_Create(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "M",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		name, _ := valueobject.NewEmployeeName(factory.FirstName, factory.MiddleName, factory.LastName, factory.NickName)
		gender, _ := enum.ParseGender(factory.Gender)
		nationality, _ := enum.ParseNationality(factory.Nationality)
		maritalStatus, _ := enum.ParseMaritalStatus(factory.MaritalStatus)
		religion, _ := enum.ParseReligion(factory.Religion)

		personalInfo, err := factory.Create()

		assert.Nil(t, err)
		assert.NotNil(t, personalInfo)
		assert.Equal(t, *name, personalInfo.Name())
		assert.False(t, personalInfo.BirthDate().IsZero())
		assert.Equal(t, factory.PlaceOfBirth, personalInfo.PlaceOfBirth())
		assert.Equal(t, gender, personalInfo.Gender())
		assert.Equal(t, nationality, personalInfo.Nationality())
		assert.Equal(t, maritalStatus, personalInfo.MaritalStatus())
		assert.Equal(t, religion, personalInfo.Religion())
	})

	t.Run("ValidInputZeroBirthDate", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Time{},
			PlaceOfBirth:  "jakarta",
			Gender:        "M",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		name, _ := valueobject.NewEmployeeName(factory.FirstName, factory.MiddleName, factory.LastName, factory.NickName)
		gender, _ := enum.ParseGender(factory.Gender)
		nationality, _ := enum.ParseNationality(factory.Nationality)
		maritalStatus, _ := enum.ParseMaritalStatus(factory.MaritalStatus)
		religion, _ := enum.ParseReligion(factory.Religion)

		personalInfo, err := factory.Create()

		assert.Nil(t, err)
		assert.NotNil(t, personalInfo)
		assert.Equal(t, *name, personalInfo.Name())
		assert.False(t, personalInfo.BirthDate().IsZero())
		assert.Equal(t, factory.PlaceOfBirth, personalInfo.PlaceOfBirth())
		assert.Equal(t, gender, personalInfo.Gender())
		assert.Equal(t, nationality, personalInfo.Nationality())
		assert.Equal(t, maritalStatus, personalInfo.MaritalStatus())
		assert.Equal(t, religion, personalInfo.Religion())
	})

	t.Run("InvalidName", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      "",
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "M",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, "last name cannot be empty")
	})

	t.Run("EmptyPlaceOfBirth", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "",
			Gender:        "M",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, "place of birth cannot be empty")
	})

	t.Run("InvalidNationality", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "M",
			Nationality:   "wno",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		_, errnationality := enum.ParseNationality(factory.Nationality)

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, errnationality.Error())
	})

	t.Run("InvalidGender", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "X",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islam",
		}

		_, errgender := enum.ParseGender(factory.Gender)

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, errgender.Error())
	})

	t.Run("InvalidMaritalStatus", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "m",
			Nationality:   "wni",
			MaritalStatus: "singlo",
			Religion:      "islam",
		}

		_, errmaritalStatus := enum.ParseMaritalStatus(factory.MaritalStatus)

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, errmaritalStatus.Error())
	})

	t.Run("InvalidReligion", func(t *testing.T) {
		factory := employee_entity.PersonalInfoFactory{
			FirstName:     faker.FirstName(),
			MiddleName:    fake.Words(1),
			NickName:      fake.Words(1),
			LastName:      faker.LastName(),
			BirthDate:     time.Now(),
			PlaceOfBirth:  "jakarta",
			Gender:        "m",
			Nationality:   "wni",
			MaritalStatus: "single",
			Religion:      "islamm",
		}

		_, errreligion := enum.ParseReligion(factory.Religion)

		personalInfo, err := factory.Create()

		assert.NotNil(t, err)
		assert.Nil(t, personalInfo)
		assert.EqualError(t, err, errreligion.Error())
	})
}
