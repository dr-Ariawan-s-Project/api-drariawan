package users

import "time"

type UsersCore struct {
	ID             int
	Name           string
	Email          string
	Password       string
	Role           string
	UrlPicture     string
	Specialization string
	State          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type UserData interface {
	Insert(data UsersCore) (UsersCore, error)
	Update(data UsersCore, id int) (UsersCore, error)
	Delete(id int) error
	// Select(search string, rp int, page int) ([]*UsersCore, error)
	// GetByID(id int) (*UsersCore, error)
	// GetByUsername(username string) (*UsersCore, error)
}

type UserService interface {
	Insert(data UsersCore) (int, error)
	Update(data UsersCore, id int) error
	Delete(id int) error
	// FindById(int) (*UsersCore, error)
	// FindByUsernameOrEmail(string) (*UsersCore, error)
	// FindAll(int, int, string) []*UsersCore
}

type UserHandler interface {
}
