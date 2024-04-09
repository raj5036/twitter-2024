package router

import (
	"github.com/gorilla/mux"
	userController "github.com/raj5036/twitter-2024/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user/registration", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/user/login", userController.LoginUser).Methods("POST")

	return router
}
