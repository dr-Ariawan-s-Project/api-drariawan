package data

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

// Insert implements users.UserData.
func (*userQuery) Insert(data users.UsersCore) (users.UsersCore, error) {
	panic("unimplemented")
}

// Update implements users.UserData.
func (*userQuery) Update(data users.UsersCore, id int) (users.UsersCore, error) {
	panic("unimplemented")
}

// Delete implements users.UserData.
func (*userQuery) Delete(id int) error {
	panic("unimplemented")
}
