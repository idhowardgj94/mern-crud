package models

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID     *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Sex    string              `bson:"sex" json:"sex"`
	Name   string              `bson:"name" json:"name"`
	Email  string              `bson:"email" json:"email"`
	Age    int                 `bson:"age" json:"age"`
	Gender string              `bson:"gender" json:"gender"`
}

func SanitizeName(name string) string {
	ret := strings.Title(strings.ToLower(name))
	return ret
}

func SanitizeEmail(email string) string {
	ret := strings.ToLower(email)
	return ret
}

func SanitizeAge(age string) (int, error) {
	val, err := strconv.Atoi(age)
	return val, err
}

func SanitizeGender(gender string) string {
	if gender == "m" || gender == "f" {
		return gender
	}
	return ""
}
