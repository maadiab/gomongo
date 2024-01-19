package router

import (
	"github.com/gorilla/mux"
	"github.com/maadiab/gomongo/controller"
	"github.com/maadiab/gomongo/database"
)

func New(db database.Database) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/lessons", controller.GetAllRecordsFromDB).Methods("GET")

	userAPI := controller.UserAPI{db: user_DB}

	router.HandleFunc("/api/lesson", userAPI.AddUser).Methods("POST")
	router.HandleFunc("/api/lesson/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/lesson/{id}", controller.DeleteOneRecordFromDB).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllRecordsFromDB).Methods("DELETE")

	return router
}
