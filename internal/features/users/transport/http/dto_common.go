package users_transport_http

import "github.com/glebateee/taskapp/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id" example:"12"`
	Version     int     `json:"version" example:"2"`
	FullName    string  `json:"full_name" example:"Номинальный Номинал Номиналович"`
	PhoneNumber *string `json:"phone_number" example:"+1234567890123"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))
	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}
	return usersDTO
}
