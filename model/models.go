package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lesson struct {
	ID      primitive.ObjectID `json:"_id",omitempty bson:"_id,omitempty"`
	Lesson  string             `json:"lessonname",omitempty`
	Watched bool               `json:"is_Watched",omitempty`
}
