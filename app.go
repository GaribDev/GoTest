package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func FindMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func CreateMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
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
