package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
	ID     *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Sex    string              `bson:"sex" json:"sex"`
	Name   string              `bson:"name" json:"name"`
	Email  string              `bson:"email" json:"email"`
	Age    int                 `bson:"age" json:"age"`
	Gender string              `bson:"gender" json:"gender"`
}

// (u *User)
