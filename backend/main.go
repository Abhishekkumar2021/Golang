package main

import (
	"backend/routes"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", routes.Router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
