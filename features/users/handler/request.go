package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/users"

type UserRequest struct {
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	Role           string `json:"role" form:"role"`
	UrlPicture     string `json:"picture" form:"picture"`
	Specialization string `json:"specialization" form:"specialization"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ReqToCore(data interface{}) *users.UsersCore {
	res := users.UsersCore{}
	switch data.(type) {
	case UserRequest:
		cnv := data.(UserRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Role = cnv.Role
		res.UrlPicture = cnv.UrlPicture
		res.Specialization = cnv.Specialization

	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
