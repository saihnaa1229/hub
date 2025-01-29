package routes

import (
	"hub/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()

	// Define all routes here
	api.HandleFunc("/users", controller.GetUsers).Methods(http.MethodGet)
	api.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	api.HandleFunc("/login", controller.LoginHandler).Methods(http.MethodPost)
	api.HandleFunc("/upload", controller.UploadVideo).Methods(http.MethodPost)
	api.HandleFunc("/video/first", controller.GetFirstVideo).Methods(http.MethodGet)

	// Video streaming route
	api.HandleFunc("/video", controller.GetVideo).Methods("GET")
}
