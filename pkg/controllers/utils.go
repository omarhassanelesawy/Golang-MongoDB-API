package controllers

import (
	"ideanesttask/pkg/database/mongodb/models"
	"os"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v2"
)

func GetDatabaseConnectionString(configFile string) string {
	// Open the configuration file
	file, err := os.Open(configFile)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Decode the YAML configuration
	var config map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return ""
	}

	// Extract the database connection string from the map
	connectionString, ok := config["database"].(map[string]interface{})["connection_string"].(string)
	if !ok {
		return ""
	}

	// Return the database connection string
	return connectionString
}

var jwtSecret = []byte("your_jwt_secret_key")

// GenerateAccessToken generates an access token for the given user
func GenerateAccessToken(user models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Sign the token with the secret key
	accessToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
