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

// SanitizeName -> string -> string
func SanitizeName(name string) string {
	ret := strings.Title(strings.ToLower(name))
	return ret
}

// SanitizeEmail -> string -> string
func SanitizeEmail(email string) string {
	ret := strings.ToLower(email)
	return ret
}

// SanitizeAge -> string -> (int, error)
func SanitizeAge(age string) (int, error) {
	val, err := strconv.Atoi(age)
	return val, err
}

// SanitizeGender -> string -> string
func SanitizeGender(gender string) string {
	if gender == "m" || gender == "f" {
		return gender
	}
	return ""
}

// Validate validate value
func (data *User) Validate() map[string]string {
	// validate
	errors := make(map[string]string)
	vname := SanitizeName(data.Name)
	vemail := SanitizeEmail(data.Email)
	if vemail == "" {
		errors["email"] = "email is required!"
	}
	if data.Age < 5 {
		errors["age"] = "you are too young to this!"
	}
	if data.Age > 130 {
		errors["ages"] = "your are too old to this!"
	}
	vgender := SanitizeGender(data.Gender)
	if vgender == "" {
		errors["gender"] = "gender is required"
	}
	if vname == "" {
		errors["name"] = "name is required"
	}
	return errors
}
