package controllers

import (
	"os"

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
