package middleware

import (
	"fmt"
	"ideanesttask/pkg/database/mongodb/models"
	"ideanesttask/pkg/database/mongodb/repository"
	"log"
)

func AddToken(token models.Token) {
	repository.ConnectDB()
	db := "admin"
	collection := "tokens"
	coll := repository.GlobalConnection.Database(db).Collection(collection)

	// Perform any necessary validation
	insertResult, err := coll.InsertOne(repository.GlobalContext, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

}
