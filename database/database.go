package database

import "go.mongodb.org/mongo-driver/mongo"

const ConnStr = "mongodb+srv://maadiab:Aa@123@cluster0.elmsjsb.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Lessons"
const colName = "LessonList"

// MOST IMPORTANT

var collection *mongo.Collection

func init() {

}
