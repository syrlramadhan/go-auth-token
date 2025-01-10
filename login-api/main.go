package main

import (
	"fmt"
	"net/http"
	"strings"
	"tugas-pemrograman-web/config"
	"tugas-pemrograman-web/controller"
	"tugas-pemrograman-web/repository"
	"tugas-pemrograman-web/service"
	"tugas-pemrograman-web/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Print("tugas pemrograman web restfull api")

	mysql, err := config.ConnectToDatabase()
	util.SentPanicIfError(err)

	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository, mysql)
	userController := controller.NewUserControllerImpl(userService)

	router := httprouter.New()

	//create
	router.POST("/api/user/create", userController.CreateUser)

	//login
	router.POST("/api/user/login", userController.LoginUser)

	//read
	router.GET("/api/user", userController.ReadUser)

	//update
	router.PUT("/api/user/update/:userId", userController.UpdateUser)

	//delete
	router.DELETE("/api/user/delete/:userId", userController.DeleteUser)

	//verify user
	router.GET("/api/user/me", VerifyJWT(userController.GetUserInfo))
	
	handler := corsMiddleware(router)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.AppPort),
		Handler: handler,
	}

	errServer := server.ListenAndServe()
	util.SentPanicIfError(errServer)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func VerifyJWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid Token Format", http.StatusUnauthorized)
			return
		}

		claims := &service.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("User-Email", claims.Email)
		next(w, r, ps)
	}
}