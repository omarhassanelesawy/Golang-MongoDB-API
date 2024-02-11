package models

import "fmt"

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

func PrintUserInfo(u User) {
	fmt.Printf("Name: %s, Email: %s\n", u.Name, u.Email)
}
