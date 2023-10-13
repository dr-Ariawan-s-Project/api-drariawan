package auth

type UserCore struct {
	Id       uint
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string
	Role     string
}

type PatientCore struct {
	Id       string
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string
	Phone    string
}

type AuthDataInterface interface {
	CreateToken(id any, role string) (token string, err error)
	GetUserByEmail(email string) (UserCore, error)
	GetPatientByEmail(email string) (PatientCore, error)
}

type AuthServiceInterface interface {
	Login(email string, password string) (string, error)
	LoginPatient(email string, password string) (string, error)
}
