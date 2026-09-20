package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := chi.NewRouter()

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "68753",
		Title: "Hansel & Gretel The Witch Hunters",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Doe",
		},
	})

	r.Get("/getMovies", getMovies)
	r.Get("/getMovie", getMovie)
	r.Post("/createMovie", createMovie)
	r.Put("/updateMovie", updateMovie)
	r.Delete("/deleteMovie", deleteMovie)

	fmt.Println("Server is running on PORT: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
