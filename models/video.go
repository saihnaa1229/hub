package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Video represents a video file in the database
type Video struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"` // MongoDB Object ID
	Title       string             `json:"title" bson:"title"`      // Video title
	Description string             `json:"description" bson:"description"` // Video description
	FileName    string             `json:"fileName" bson:"fileName"` // File name in GridFS
	UploadDate  string             `json:"uploadDate" bson:"uploadDate"` // Upload date
}
