package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/maadiab/gomongo/database"
	"github.com/maadiab/gomongo/router"
)

func main() {
	db, err := database.New(context.Background())
	if err != nil {
		log.Fatalln("could not connect to database:", err)
	}

	fmt.Println("Server is getting started...")
	r := router.New(db)
	log.Println("listening at port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
