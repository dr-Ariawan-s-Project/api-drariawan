package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type UsersCore struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Password       string     `json:"password"`
	Role           string     `json:"role"`
	UrlPicture     string     `json:"picture"`
	Specialization string     `json:"specialization"`
	State          string     `json:"state"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type UserData interface {
	Insert(data UsersCore) (UsersCore, error)
	Update(data UsersCore, id int) error
	Delete(id int) error
	FindAll(search string, rp int, page int) ([]UsersCore, error)
	FindByID(id int) (UsersCore, error)
	FindByUsernameOrEmail(username string) (UsersCore, error)
}

type UserService interface {
	Insert(data UsersCore) (UsersCore, error)
	Update(data UsersCore, token interface{}) error
	Delete(id int) error
	FindAll(search string, rp, page int) ([]UsersCore, error)
	FindById(id int) (UsersCore, error)
	FindByUsernameOrEmail(username string) (UsersCore, error)
}

type UserHandler interface {
	Insert() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	FindAll() echo.HandlerFunc
	FindById() echo.HandlerFunc
	FindByUsernameOrEmail() echo.HandlerFunc
}
