package util

import (
	"tugas-pemrograman-web/dto"
	"tugas-pemrograman-web/model"
)

func ToUserResponse(user model.User) dto.UserResponse {
	return dto.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Pass:  user.Password,
	}
}

func ToUserListResponse(user []model.User) []dto.UserResponse {
	var userResponses []dto.UserResponse
	for _, users := range user {
		userResponses = append(userResponses, ToUserResponse(users))
	}
	return userResponses
}
