package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lesson struct {
	ID      primitive.ObjectID `json: "_id",omitempty bson: "_id",omitempty`
	Lesson  string             `json: "lesson_name",omitempty`
	Watched bool               `json: "watched",omitempty`
}
