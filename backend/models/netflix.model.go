package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Year    int                `json:"year,omitempty" bson:"year,omitempty"`
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
