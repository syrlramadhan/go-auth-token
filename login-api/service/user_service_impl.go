package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"tugas-pemrograman-web/dto"
	"tugas-pemrograman-web/model"
	"tugas-pemrograman-web/repository"
	"tugas-pemrograman-web/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte("secret_key")

type userServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *userServiceImpl) CreateUser(ctx context.Context, userRequest dto.CreateUserRequest) dto.UserResponse {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	user := model.User{
		Id:       uuid.New().String(),
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Pass,
	}

	createUser, errSave := service.UserRepository.CreateUser(ctx, tx, user)
	util.SentPanicIfError(errSave)

	return convertToResponseDTO(createUser)
}

func convertToResponseDTO(user model.User) dto.UserResponse {
	return dto.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Pass:  user.Password,
	}
}

func (service *userServiceImpl) ReadUser(ctx context.Context) []dto.UserResponse {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	user := service.UserRepository.ReadUser(ctx, tx)

	return util.ToUserListResponse(user)
}

func (service *userServiceImpl) UpdateUser(ctx context.Context, userRequest dto.UpdateUserRequest, idUser string) dto.UserResponse {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindById(ctx, tx, idUser)
	util.SentPanicIfError(err)

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Pass

	user = service.UserRepository.UpdateUser(ctx, tx, user)

	return util.ToUserResponse(user)
}

func (service *userServiceImpl) DeleteUser(ctx context.Context, idUser string) dto.UserResponse {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)
	defer util.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindById(ctx, tx, idUser)
	util.SentPanicIfError(err)

	deletedUser := service.UserRepository.DeleteUser(ctx, tx, user)

	return util.ToUserResponse(deletedUser)
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (service *userServiceImpl) GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "go-auth-example",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func (service *userServiceImpl) LoginUser(ctx context.Context, loginRequest dto.LoginUserRequest) (string, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %v", err)
	}
	defer util.CommitOrRollBack(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, loginRequest.Email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if user.Password != loginRequest.Pass {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := service.GenerateJWT(loginRequest.Email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (service *userServiceImpl) GetUserInfoByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	tx, err := service.DB.Begin()
	util.SentPanicIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("user not found")
	}

	userResponse := dto.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Pass:  user.Password,
	}

	return userResponse, nil
}