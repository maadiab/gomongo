package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maadiab/gomongo/helper"
	models "github.com/maadiab/gomongo/model"
)

// Get all records
func GetAllRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/X-www-form-urlencode")
	allRecords := helper.GetAllRecord()
	json.NewEncoder(w).Encode(allRecords)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-form-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var lesson models.Lesson
	_ = json.NewDecoder(r.Body).Decode(&lesson)
	helper.InsertOneRecord(lesson)
	json.NewEncoder(w).Encode(lesson)

}

// Mark as watched

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-form-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)

	helper.UpdateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params)

}

// Delete one record

func DeleteOneRecordFromDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-form-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)

	helper.DeleteOneRecord(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

// Delete all lessons

func DeleteAllRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-form-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := helper.DeleteAllRecords()

	json.NewEncoder(w).Encode(count)

}
