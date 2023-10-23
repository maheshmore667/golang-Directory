package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maheshmore667/mongoapi/router"
)

func main() {
	fmt.Println("Mongo Learning")
	r := router.Router()
	fmt.Println("App is starting ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("App is listening on port 4000")
}
