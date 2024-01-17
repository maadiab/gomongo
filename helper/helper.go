package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/maadiab/gomongo/database"
	Models "github.com/maadiab/gomongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Helper

func InsertOneRecord(lesson Models.Lesson) {
	inserted, err := database.Mcollection.InsertOne(context.Background(), lesson)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 lesson in db with id : ", inserted.InsertedID)
}

// update 1 record

func UpdateOneRecord(lessonId string) {
	id, _ := primitive.ObjectIDFromHex(lessonId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := database.Mcollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record

func DeleteOneRecord(lessonId string) {
	id, _ := primitive.ObjectIDFromHex(lessonId)
	filter := bson.M{"_id": id}
	deleteCount, err := database.Mcollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("lesson got deleted with delete count: ", deleteCount)

}

// delete all or many

func DeleteAllRecords() int64 {
	deleteResult, err := database.Mcollection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of lessons deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

// All lessons from mongoDB

func GetAllRecord() []primitive.M {
	cur, err := database.Mcollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var lessons []primitive.M

	for cur.Next(context.Background()) {
		var lesson bson.M
		err := cur.Decode(&lesson)
		if err != nil {
			log.Fatal(err)
		}
		lessons = append(lessons, lesson)

	}

	defer cur.Close(context.Background())
	return lessons
}
