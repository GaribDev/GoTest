package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/GaribDev/GoTest/dao"
	. "github.com/GaribDev/GoTest/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = MoviesDAO{}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func FindMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func CreateMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func DeleteMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMoviesEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMoviesEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMoviesEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMoviesEndPoint).Methods("GET")
	if err := http.ListenAndServe(":27017", r); err != nil {
		log.Fatal(err)
	}
}
