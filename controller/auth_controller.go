package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"hub/config"
	"hub/models"

	"go.mongodb.org/mongo-driver/bson"
)
// LoginHandler handles user login by username and password
// @Summary Login User
// @Description Authenticate a user by username and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginCredentials true "User Credentials"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Invalid username or password"
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var loginData models.LoginCredentials

    // Parse the request body
    if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Connect to the users collection
    collection := config.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Find the user by username and password
    var user models.User
    filter := bson.M{"username": loginData.Username, "password": loginData.Password}
    err := collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    // Return the user data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
