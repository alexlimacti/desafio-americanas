package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name" validate:"required"`
	Climate string             `json:"climate" validate:"required"`
	Ground  string             `json:"ground" validate:"required"`
}
