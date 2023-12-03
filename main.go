package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In this case `json: "id"` doesn't work because of the
// space after ':'. It should be `json:"id"`
type Movie struct {
	ID    string `json:"id"`
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
	//Why Director is a pointer here? Look at readme (4)
	Director *Director `json:"director"`
}

type Director struct {
	Firstname  string `json:"firstname"`
	Secondname string `json:"secondname"`
}

// GET ALL MOVIES
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// DELETE A MOVIE
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//To DO
}

var movies []Movie

func main() {
	//mux.NewRouter() is imported from Gorilla Mux
	router := mux.NewRouter()

	movies = append(movies,
		Movie{
			ID:    "1",
			ISBN:  "34232",
			Title: "Lele",
			//Why Director is an '&' here? Look at readme (4)
			Director: &Director{
				Firstname:  "Rafid",
				Secondname: "Niloy",
			}})

	movies = append(movies,
		Movie{
			ID:    "2",
			ISBN:  "94724",
			Title: "Lele 2",
			Director: &Director{
				Firstname:  "Farhan",
				Secondname: "Bhotka",
			}})

	movies = append(movies,
		Movie{
			ID:    "3",
			ISBN:  "73842",
			Title: "Journey of a Magi: Farhan",
			Director: &Director{
				Firstname:  "Mahadi",
				Secondname: "Nikka",
			}})

	//router.HandleFunc( path, handler).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port 8000...")

	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal(err)
	}

}
