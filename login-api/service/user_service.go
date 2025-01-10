package service

import (
	"context"
	"tugas-pemrograman-web/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, userRequest dto.CreateUserRequest) dto.UserResponse
	GenerateJWT(userId string) (string, error)
	LoginUser(ctx context.Context, loginRiquest dto.LoginUserRequest) (string, error)
	ReadUser(ctx context.Context) []dto.UserResponse
	UpdateUser(ctx context.Context, userRequest dto.UpdateUserRequest, idUser string) dto.UserResponse
	DeleteUser(ctx context.Context, idUser string) dto.UserResponse
	GetUserInfoByEmail(ctx context.Context, email string) (dto.UserResponse, error)
}
