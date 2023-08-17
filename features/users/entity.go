package users

import "time"

type Users struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Role       string
	Picture    string
	Status     string
	Created_at time.Time
	Updated_at time.Time
}

type UserRepositories interface {
	Insert(data Users) (Users, error)
	Update(data Users, id int) (Users, error)
	Delete(id int) error
	Select(search string, rp int, page int) ([]Users, error)
	GetByID(id int) (Users, error)
	GetByUsername(username string) (Users, error)
}
