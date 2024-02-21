package controllers

import (
	"backend/db"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Send a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var data = map[string]interface{}{
		"status":  "success",
		"message": "Welcome to Netflix API",
	}
	json.NewEncoder(w).Encode(data)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	// read the JSON data from the request body
	// decode the JSON data
	// insert the data into the database
	// send a JSON response

	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)

	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Invalid request",
		})
		return
	}

	// insert the data into the database
	result, err := db.Collection.InsertOne(context.Background(), movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "success",
		"message":     "Movie added successfully",
		"movie":       movie,
		"inserted_id": result.InsertedID,
	})
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// get all the movies from the database
	// send a JSON response

	var movies []models.Netflix
	cursor, err := db.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie models.Netflix
		cursor.Decode(&movie)
		fmt.Printf("Movie: %+v\n", movie)
		movies = append(movies, movie)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"movies": movies,
	})
}
