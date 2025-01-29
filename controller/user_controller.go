package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"hub/config"
	"hub/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary Get Users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := collection.Find(ctx, bson.M{}, options.Find())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// @Summary Create User
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	collection := config.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Set the creation timestamp
	user.CreatedAt = time.Now()

	// Insert the user into the database
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Convert the inserted ID to a string (if using MongoDB's ObjectID)
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid.Hex()
	} else {
		http.Error(w, "Failed to retrieve inserted ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
