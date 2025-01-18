package models

// LoginCredentials represents the username and password required for login
type LoginCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
