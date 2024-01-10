package patient

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
)

type Core struct {
	ID             string
	Name           string
	Email          string
	Password       string
	NIK            string
	DOB            *time.Time
	Phone          string
	Gender         *string
	MarriageStatus string
	Nationality    string
	PartnerID      *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Partner        *Core
}

type PatientDataInterface interface {
	Insert(data Core) (*Core, error)
	Update(id string, data Core) (*Core, error)
	Delete(id string) error
	Select(search string, offset int, limit int) ([]Core, error)
	SelectById(id string) (*Core, error)
	SelectByEmailOrPhone(str string) (*Core, error)
	CheckByEmailAndPhone(email string, phone string) (*Core, error)
	CountPartner(partnerId string) (int, error)
	CountAllPatient() (int, error)
	SelectAllNIK() ([]string, error)
	CountByFilter(search string) (int64, error)
}

type PatientServiceInterface interface {
	Insert(data Core, partnerEmail string) (*Core, error)
	Update(data Core) (*Core, error)
	Delete(id string) error
	FindAll(search string, page int, perPage int) ([]Core, error)
	FindById(id string) (*Core, error)
	CheckByEmailAndPhone(email string, phone string) (*Core, error)
	CountPartner(partnerId string) (int, error)
	CountAllPatient() (int, error)
	GetPagination(search string, page int, perPage int) (helpers.Pagination, error)
}
