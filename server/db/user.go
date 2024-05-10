package db

import (
	"context"
	"meeting_scheduler_api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user model.User) (primitive.ObjectID,error) {
	res,err:=database.Collection("user").InsertOne(context.Background(),user)
	if err!=nil {
		return primitive.NilObjectID,err;
	}
	return res.InsertedID.(primitive.ObjectID),nil
}