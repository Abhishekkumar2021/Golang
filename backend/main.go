package main

import (
	"backend/routes"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes.Router))
}
