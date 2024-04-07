package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	PhoneNumber string             `json:"phoneNumber,omitempty"`
	DOB         string             `json:"dob,omitempty"`
	Tweets      any                `json:"tweets,omitempty"`
	Likes       any                `json:"likes,omitempty"`
	Replies     any                `json:"replies,omitempty"`
}
