package users

import "time"

type Users struct {
	ID             int
	FullName       string
	Email          string
	Password       string
	Role           string
	Picture        string
	Specialization string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Repositories interface {
	Insert(data Users) (*Users, error)
	Update(data Users, id int) (*Users, error)
	Delete(id int) error
	Select(search string, rp int, page int) ([]*Users, error)
	GetByID(id int) (*Users, error)
	GetByUsername(username string) (*Users, error)
}
