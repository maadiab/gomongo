package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnStr = "mongodb+srv://maadiab:Aa123123@cluster0.elmsjsb.mongodb.net/?retryWrites=true&w=majority"

// Aa123123
const dbName = "Lessons"
const colName = "LessonList"

// MOST IMPORTANT

var Mcollection *mongo.Collection

func init() {
	// Client options
	clientOptions := options.Client().ApplyURI(ConnStr)

	// Connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection sucess!!")

	Mcollection = client.Database(dbName).Collection(colName)

	// Collection reference is  ready

	fmt.Println("Collection instance is ready!")

}

// Mongo
