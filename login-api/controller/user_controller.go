package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ReadUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUserInfo(writer http.ResponseWriter, request *http.Request, paramas httprouter.Params)
}
