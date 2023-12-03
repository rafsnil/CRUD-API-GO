package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]

	//looping to find the movie with the desire id
	for index, movie := range movies {
		if movie.ID == movieId {
			// ... is used to spread the slices into individual elements
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

// GET MOVIE WITH ID HANDLER
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]

	for _, movie := range movies {

		if movie.ID == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

// CREATE MOVIE HANDLER
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	/*Decoding response body to json and using pointer to save
	the given data to newMovie struct.*/
	json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(100000))

	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(movies)
}

// UPDATE MOVIE HANDLER
/*
PSEUDO CODE:
1. set content type as json
2. get movie id from params
3. loop over movies and get desired movie x
4. delete movie x
5. add new movie with x.id that will be sent through postman
*/
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//seting content type as json
	w.Header().Set("Content-Type", "application/json")

	//getting movie id from params
	params := mux.Vars(r)
	upMovieId := params["id"]

	//Iterating movies to find desired movie
	for index, movie := range movies {
		if movie.ID == upMovieId {
			//Deleting movie after finding it using a cheeky method
			movies = append(movies[:index], movies[index+1:]...)

			//Creating a new movie
			var newMovie Movie
			//Converting the response body to json and putting the info in newMovie
			json.NewDecoder(r.Body).Decode(&newMovie)
			//Getting data for newMovie
			newMovie.ID = upMovieId
			newMovie.Title = params["title"]
			//Putting the newMovie back to the slice
			movies = append(movies, newMovie)
			//Sending back response
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

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
				Firstname:  "Niloy",
				Secondname: "Baitta",
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
				Secondname: "Kala",
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
