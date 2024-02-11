package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GlobalConnection *mongo.Client
var GlobalContext context.Context
var Temp string

func ConnectDB() {
	localconnection, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/admin"))
	if err != nil {
		log.Fatal("First")
	}
	GlobalConnection = localconnection
	localcontext, _ := context.WithTimeout(context.Background(), 100*time.Second)
	GlobalContext = localcontext
	err = GlobalConnection.Connect(GlobalContext)
	if err != nil {
		log.Fatal("Second")
	}
	// defer GlobalConnection.Disconnect(GlobalContext)
}
