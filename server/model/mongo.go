package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"login" json:"login"`
	Password string             `bson:"password" json:"possword"`
	User     primitive.ObjectID `bson:"user,omitempty" json:"user,omitempty"`
}

type User struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name            string               `bson:"name" json:"name"`
	Age             int                  `bson:"age" json:"age"`
	Gender          string               `bson:"gender" json:"gender"`
	Address         string               `bson:"address" json:"address"`
	ScheduleCount   int                  `bson:"schedule_cout_per_day" json:"schedule_cout_per_day"`
	Role            primitive.ObjectID   `bson:"role,omitempty" json:"role,omitempty"`
	Schedule        []primitive.ObjectID `bson:"schedule" json:"schedule"`
	ScheduleRequest []primitive.ObjectID `bson:"schedule_request" json:"schedule_request"`
}

type Role struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Role string             `bson:"role" json:"role"`
}

type Schedule struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	ScheduledBy primitive.ObjectID `bson:"scheduled_by" json:"scheduled_by"`
	ScheduledTo primitive.ObjectID `bson:"scheduled_to" json:"scheduled_to"`
	Comment     string             `bson:"comment" json:"comment"`
	Status      string             `bson:"status" json:"status"`
}
