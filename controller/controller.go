package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maadiab/gomongo/core"
	"github.com/maadiab/gomongo/database"
	"github.com/maadiab/gomongo/helper"
)

type UserAPI struct {
	user_DB database.UserDB
}

// Get all records
func GetAllRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/X-www-form-urlencode")
	allRecords := helper.GetAllRecord()
	json.NewEncoder(w).Encode(allRecords)
}

func (a UserAPI) AddUser(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/x-form-www-form-urlencode")
	//w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// handle error
	}

	user, err = a.user_DB.AddUser(r.Context(), user)
	if err != nil {
		// handle error
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		// handler error
	}
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
