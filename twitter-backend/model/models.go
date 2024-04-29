package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	PhoneNumber string             `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	DOB         string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Tweets      any                `json:"tweets,omitempty" bson:"tweets,omitempty"`
	Likes       any                `json:"likes,omitempty" bson:"likes,omitempty"`
	Replies     any                `json:"replies,omitempty" bson:"replies,omitempty"`
}
