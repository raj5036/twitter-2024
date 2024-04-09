package router

import (
	"net/http"

	"github.com/gorilla/mux"
	userController "github.com/raj5036/twitter-2024/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user/registration", userController.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/user/login", userController.LoginUser).Methods(http.MethodPost)

	return router
}
