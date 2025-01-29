package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
