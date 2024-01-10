package users

import (
	"time"
)

type UsersCore struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	Password       string     `json:"password"`
	Role           string     `json:"role"`
	UrlPicture     string     `json:"picture"`
	Specialization string     `json:"specialization"`
	State          string     `json:"state"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type Service interface {
	Insert(data UsersCore, role string) (UsersCore, error)
	Update(data UsersCore, id int) error
	Delete(id int) error
	FindAll(search string, rp, page int) ([]UsersCore, error)
	FindById(id int) (UsersCore, error)
	GetPagination(search string, page int, perPage int) (map[string]any, error)
}

type Data interface {
	Insert(data UsersCore) (UsersCore, error)
	Update(data UsersCore, id int) error
	Delete(id int) error
	FindAll(search string, rp int, page int) ([]UsersCore, error)
	FindByID(id int) (UsersCore, error)
	CountByFilter(search string) (int64, error)
}
