package models

import (
	"fmt"
	"time"
)

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type ResponseSchema struct {
	Message string `bson:"message"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Organization struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Token struct {
	Value      string    `json:"value" bson:"value"`
	ExpiryTime time.Time `json:"expiry_time" bson:"expiry_time"`
}

func PrintUserInfo(u User) {
	fmt.Printf("Name: %s, Email: %s\n", u.Name, u.Email)
}
