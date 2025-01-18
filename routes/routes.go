package routes

import (
	"hub/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()

	// Define all routes here
	api.HandleFunc("/users", controller.GetUsers).Methods("GET")
	api.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	api.HandleFunc("/upload", controller.UploadVideo).Methods("POST")
	api.HandleFunc("/video/first",controller.GetFirstVideo).Methods("GET")

	// Video streaming route
	api.HandleFunc("/video", controller.GetVideo).Methods("GET")
}
