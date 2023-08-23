package auth

type UserCore struct {
	Id       uint
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string
	Role     string
}

type AuthDataInterface interface {
	CreateToken(id int, role string) (token string, err error)
	GetUserByEmail(email string) (UserCore, error)
}

type AuthServiceInterface interface {
	Login(email string, password string) (string, error)
}
