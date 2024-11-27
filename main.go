package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Simulates a database
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) // Enconding the response
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // returns all the other movies
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Geting access to all the params that client requested
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // decodes the json body
	movie.ID = strconv.Itoa((rand.IntN(1_000_000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // returns the new movie
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// First: Delete the movie
	// Second: Add the new movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	port := "localhost:8000"

	movies = append(movies, Movie{ID: "1", Isbn: "4653233", Title: "Movie One", Director: &Director{FirstName: "Eichi", LastName: "Arakaki"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4653234", Title: "Movie Two", Director: &Director{FirstName: "Sayuri", LastName: "Arakaki"}})
	movies = append(movies, Movie{ID: "3", Isbn: "4653235", Title: "Movie Three", Director: &Director{FirstName: "Ayumi", LastName: "Arakaki"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("STARTING SERVER AT %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
