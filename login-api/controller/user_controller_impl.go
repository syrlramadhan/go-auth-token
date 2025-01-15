package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"tugas-pemrograman-web/dto"
	"tugas-pemrograman-web/service"
	"tugas-pemrograman-web/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type userControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (controller *userControllerImpl) CreateUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	requestCreate := dto.CreateUserRequest{}
	util.ReadFromRequestBody(request, &requestCreate)

	responseDTO := controller.UserService.CreateUser(request.Context(), requestCreate)
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "create user successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}
	
func (controller *userControllerImpl) ReadUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	responseDTO := controller.UserService.ReadUser(request.Context())
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "read user successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestUpdate := dto.UpdateUserRequest{}
	util.ReadFromRequestBody(request, &requestUpdate)

	userId := params.ByName("userId")

	responseDTO := controller.UserService.UpdateUser(request.Context(), requestUpdate, userId)
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "update user successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) UpdatePhoto(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20) // Max 10MB
	util.SentPanicIfError(err)

	file, handler, err := request.FormFile("photo") // "photo" harus sama dengan key pada form data
	if err != nil {
		http.Error(writer, "Failed to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	userId := params.ByName("userId")
	user, err := controller.UserService.FindById(request.Context(), userId)
	util.SentPanicIfError(err)

	fileName := fmt.Sprintf("%s.jpeg", user.Email)

	handler.Filename = fileName
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filePath := filepath.Join(uploadDir, handler.Filename)
	out, err := os.Create(filePath)
	util.SentPanicIfError(err)
	defer out.Close()

	_, err = io.Copy(out, file)
	util.SentPanicIfError(err)

	requestUpdate := dto.UpdatePhotoRequest{
		Photo: handler.Filename,
	}

	responseDTO := controller.UserService.UpdatePhoto(request.Context(), requestUpdate, userId)
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Data:    responseDTO,
		Message: "update photo successfully",
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	responseDTO := controller.UserService.DeleteUser(request.Context(), userId)

	message := fmt.Sprint("user ", responseDTO.Name, " delete successfully")
	response := dto.ResponseList{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: message,
	}

	writer.Header().Add("Content-Type", "application/json")
	util.WriteToResponseBody(writer, response)
}

func (controller *userControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginRequest dto.LoginUserRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := controller.UserService.LoginUser(r.Context(), loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := dto.ResponseToken{
		Code:    http.StatusOK,
		Status:  "OK",
		Token:   token,
		Message: "token generate successfully",
	}

	w.Header().Set("Content-.Type", "application/json")
	util.WriteToResponseBody(w, response)
}

func (controller *userControllerImpl) GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
		return
	}

	tokenString := authHeader[7:]
	claims := &service.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
		return
	}

	email := claims.Email

	userResponse, err := controller.UserService.GetUserInfoByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := dto.ResponseList {
		Code: http.StatusOK,
		Status: "OK",
		Data: userResponse,
		Message: "success login to user",
	}

	w.Header().Set("Content-Type", "application/json")
	util.WriteToResponseBody(w, response)
}
