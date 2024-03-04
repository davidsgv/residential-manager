package data

import (
	"residential-manager/internal/domain/entities"

	"github.com/google/uuid"
)

type Apartment struct {
	Block  string `json:"block"`
	Number string `json:"number"`
}

type UserResponse struct {
	Id        uuid.UUID  `json:"id"`
	Rol       string     `json:"rol"`
	Mail      string     `json:"mail"`
	Apartment *Apartment `json:"aparment,omitempty"`
}

type CreateUser struct {
	Mail      string    `json:"mail"`
	Apartment Apartment `json:"apartment"`
	Rol       string    `json:"rol"`
}

type UpdateUser struct {
	Apartment *Apartment `json:"apartment,omitempty"`
	Rol       *string    `json:"rol,omitempty"`
}

type VerifyUser struct {
	Password string `json:"password"`
}

func UserToApi(user *entities.User) *UserResponse {
	if user == nil {
		return nil
	}

	res := &UserResponse{
		Id:   user.Id,
		Rol:  *user.Rol,
		Mail: user.Mail,
	}

	if user.Apartment == nil {
		return res
	}

	res.Apartment = &Apartment{
		Block:  user.Apartment.Block,
		Number: user.Apartment.Number,
	}
	return res
}

func UsersToApi(users []entities.User) []UserResponse {
	if users == nil {
		return nil
	}

	apiUsers := []UserResponse{}
	for _, e := range users {
		apiUser := UserToApi(&e)
		apiUsers = append(apiUsers, *apiUser)
	}
	return apiUsers
}
