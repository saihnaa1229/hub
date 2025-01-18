// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Authenticate a user by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "User Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Invalid username or password",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "Uploads a video file to MongoDB using GridFS",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Videos"
                ],
                "summary": "Upload a video",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Video file to upload",
                        "name": "video",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Video uploaded successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Unable to read video file",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to upload video",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/video/first": {
            "get": {
                "description": "Streams the first video file from MongoDB GridFS",
                "produces": [
                    "video/mp4"
                ],
                "tags": [
                    "Videos"
                ],
                "summary": "Stream the first video",
                "responses": {
                    "200": {
                        "description": "Video streamed successfully",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "404": {
                        "description": "No video found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to stream video",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/video/{id}": {
            "get": {
                "description": "Streams a video file from MongoDB by its ID",
                "produces": [
                    "video/mp4"
                ],
                "tags": [
                    "Videos"
                ],
                "summary": "Stream a video",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Video streamed successfully",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "404": {
                        "description": "Video not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to stream video",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LoginCredentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Hub API",
	Description:      "API documentation for the Hub project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
