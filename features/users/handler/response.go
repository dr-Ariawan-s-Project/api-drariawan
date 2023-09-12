package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/users"

type UserResponse struct {
	Name           string `json:"name" `
	Email          string `json:"email"`
	Role           string `json:"role"`
	UrlPicture     string `json:"picture"`
	Specialization string `json:"specialization"`
}

func CoreToResponse(core users.UsersCore) UserResponse {
	return UserResponse{
		Name:           core.Name,
		Email:          core.Email,
		Role:           core.Role,
		UrlPicture:     core.UrlPicture,
		Specialization: core.Specialization,
	}
}

func CoreToResponseArray(data []users.UsersCore) []UserResponse {
	result := []UserResponse{}
	for _, data := range data {
		result = append(result, CoreToResponse(data))
	}
	return result
}
