package router

import (
	"github.com/gorilla/mux"
	"github.com/maadiab/gomongo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/lessons", controller.GetAllRecordsFromDB).Methods("GET")

	router.HandleFunc("/api/lesson", controller.CreateRecord).Methods("POST")
	router.HandleFunc("/api/lesson/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/lesson/{id}", controller.DeleteOneRecordFromDB).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controller.DeleteAllRecordsFromDB).Methods("DELETE")

	return router
}
