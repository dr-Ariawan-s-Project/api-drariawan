package patient

import "time"

type Core struct {
	ID             string
	Name           string
	Email          string
	Password       string
	NIK            string
	DOB            time.Time
	Phone          string
	Gender         string
	MarriageStatus string
	Nationality    string
	PartnerID      *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type PatientDataInterface interface {
	Insert(data Core) error
	Update(data Core) (*Core, error)
	Delete(id string) error
	Select(search string, page int, perPage int) ([]Core, error)
	SelectById(id string) (*Core, error)
	CheckByEmailOrPhone(email string, phone string) (*Core, error)
}

type PatientServiceInterface interface {
	Insert(data Core) error
	Update(data Core) (*Core, error)
	Delete(id string) error
	FindAll(search string, page int, perPage int) ([]Core, error)
	FindById(id string) (*Core, error)
	CheckByEmailOrPhone(email string, phone string) (*Core, error)
}
