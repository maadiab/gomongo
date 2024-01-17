package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maadiab/gomongo/router"
)

func main() {
	fmt.Println("MongoDB api")
	fmt.Println("Server is getting started...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000")
}
