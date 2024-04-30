package router

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/raj5036/twitter-2024/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/user/registration", controller.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/user/login", controller.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/user/delete-all", controller.DeleteAllUser).Methods(http.MethodDelete)

	// Tweet routes
	router.HandleFunc("/tweet/create", controller.CreateTweet).Methods(http.MethodPost)
	router.HandleFunc("/tweet/get/:userId", controller.GetTweetsByUserId).Methods(http.MethodGet)
	router.HandleFunc("/tweet/update/:id", controller.UpdateTweet).Methods(http.MethodPatch)
	router.HandleFunc("/tweet/delete/:id", controller.DeleteTweet).Methods(http.MethodDelete)

	return router
}
