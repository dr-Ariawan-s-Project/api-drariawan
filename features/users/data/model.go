package data

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
)

type Users struct {
	ID             int
	Name           string
	Email          string
	Phone          string
	Password       string
	Role           string
	UrlPicture     string
	Specialization string
	State          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func DataToCore(data Users) users.UsersCore {
	return users.UsersCore{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Password:       data.Password,
		Role:           data.Role,
		UrlPicture:     data.UrlPicture,
		Specialization: data.Specialization,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt,
	}
}

func CoreToData(data users.UsersCore) Users {
	return Users{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Phone:          data.Phone,
		Password:       data.Password,
		Role:           data.Role,
		UrlPicture:     data.UrlPicture,
		Specialization: data.Specialization,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt,
	}
}

func DataToCoreArray(data []Users) []users.UsersCore {
	res := []users.UsersCore{}
	for _, val := range data {
		res = append(res, DataToCore(val))
	}
	return res
}
