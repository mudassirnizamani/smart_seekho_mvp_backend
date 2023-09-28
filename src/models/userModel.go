package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	PhoneNumber  *string            `bson:"phone_number" json:"phoneNumber"`
	Name         string             `bson:"name" json:"name"`
	ProfilePic   string             `bson:"profile_pic" json:"profilePic"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_token" json:"refreshToken"`
	UserId       string             `bson:"user_id" json:"userId"`
}

type Otp struct {
	ID          primitive.ObjectID `bson:"_id"`
	PhoneNumber string             `bson:"phone_number" json:"phoneNumber"`
	Otp         string             `bson:"otp" json:"otp"`
}
