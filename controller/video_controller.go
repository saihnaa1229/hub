package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
}

// UploadVideo handles video uploads to MongoDB
// @Summary Upload a video
// @Description Uploads a video file to MongoDB using GridFS
// @Tags Videos
// @Accept multipart/form-data
// @Param video formData file true "Video file to upload"
// @Produce json
// @Success 200 {string} string "Video uploaded successfully"
// @Failure 400 {string} string "Unable to read video file"
// @Failure 500 {string} string "Unable to upload video"
// @Router /upload [post]
func UploadVideo(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Unable to read video file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Open the GridFS bucket for the "mydatabase" database
	bucket, err := gridfs.NewBucket(client.Database("mydatabase"), options.GridFSBucket().SetName("video"))
	if err != nil {
		http.Error(w, "Unable to create GridFS bucket", http.StatusInternalServerError)
		return
	}

	// Upload the video file to GridFS
	uploadStream, err := bucket.OpenUploadStream(header.Filename)
	if err != nil {
		http.Error(w, "Unable to upload video", http.StatusInternalServerError)
		return
	}
	defer uploadStream.Close()

	_, err = io.Copy(uploadStream, file)
	if err != nil {
		http.Error(w, "Failed to save video", http.StatusInternalServerError)
		return
	}

	// Respond with the file ID for reference
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Video uploaded successfully. File ID: %v", uploadStream.FileID)
}

// GetVideo streams the video by its ID
// @Summary Stream a video
// @Description Streams a video file from MongoDB by its ID
// @Tags Videos
// @Param id path string true "Video ID"
// @Produce video/mp4
// @Success 200 {file} file "Video streamed successfully"
// @Failure 404 {string} string "Video not found"
// @Failure 500 {string} string "Failed to stream video"
// @Router /video/{id} [get]
func GetVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoID := vars["id"]

	// Convert the ID string to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(videoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid video ID: %v", err), http.StatusBadRequest)
		return
	}

	// Open the GridFS bucket for the "mydatabase" database
	bucket, err := gridfs.NewBucket(client.Database("mydatabase"), options.GridFSBucket().SetName("video"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create GridFS bucket: %v", err), http.StatusInternalServerError)
		return
	}

	// Stream the video from GridFS
	downloadStream, err := bucket.OpenDownloadStream(objectID)
	if err != nil {
		// Enhance error logging for clarity
		http.Error(w, fmt.Sprintf("Video not found or failed to open stream: %v", err), http.StatusNotFound)
		return
	}
	defer downloadStream.Close()

	// Stream the video to the client
	w.Header().Set("Content-Type", "video/mp4")
	_, err = io.Copy(w, downloadStream)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to stream video: %v", err), http.StatusInternalServerError)
		return
	}
}

// GetFirstVideo streams the first video in the MongoDB GridFS bucket
// @Summary Stream the first video
// @Description Streams the first video file from MongoDB GridFS
// @Tags Videos
// @Produce video/mp4
// @Success 200 {file} file "Video streamed successfully"
// @Failure 404 {string} string "No video found"
// @Failure 500 {string} string "Failed to stream video"
// @Router /video/first [get]
func GetFirstVideo(w http.ResponseWriter, r *http.Request) {
	// Access the GridFS bucket for the "mydatabase" database
	bucket, err := gridfs.NewBucket(client.Database("mydatabase"), options.GridFSBucket().SetName("video"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create GridFS bucket: %v", err), http.StatusInternalServerError)
		return
	}

	// Query the GridFS metadata collection for the first video
	filter := bson.M{} // Empty filter to get any document
	metadataCollection := client.Database("mydatabase").Collection("video.files")
	var firstFile bson.M
	err = metadataCollection.FindOne(context.TODO(), filter).Decode(&firstFile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "No video found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Failed to find video: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Extract the ObjectID of the first video
	objectID, ok := firstFile["_id"].(primitive.ObjectID)
	if !ok {
		http.Error(w, "Invalid video ObjectID format", http.StatusInternalServerError)
		return
	}

	// Stream the video from GridFS
	downloadStream, err := bucket.OpenDownloadStream(objectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open video stream: %v", err), http.StatusInternalServerError)
		return
	}
	defer downloadStream.Close()

	// Set headers and stream the video
	w.Header().Set("Content-Type", "video/mp4")
	_, err = io.Copy(w, downloadStream)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to stream video: %v", err), http.StatusInternalServerError)
		return
	}
}
