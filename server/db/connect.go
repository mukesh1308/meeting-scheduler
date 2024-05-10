package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client;
var database *mongo.Database;

func init() {
	var err error;
	client,err=mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017"));
	if err!=nil {
		panic("database not connected");
	}
	database=client.Database("meeting_scheduler");
	fmt.Println("Database connected")
}