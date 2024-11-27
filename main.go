package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

func main() {
	r := mux.NewRouter()
	port := "localhost:8000"

	movies = append(movies, Movie{ID: "1", Isbn: "4653233", Title: "Movie One", Director: &Director{FirstName: "Eichi", LastName: "Arakaki"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4653234", Title: "Movie Two", Director: &Director{FirstName: "Sayuri", LastName: "Arakaki"}})
	movies = append(movies, Movie{ID: "3", Isbn: "4653235", Title: "Movie Three", Director: &Director{FirstName: "Ayumi", LastName: "Arakaki"}})

	r.HandleFunc("/", serveTemplate).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Servir archivos estaticos
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Printf("STARTING SERVER AT %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))

}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	// En el contexto de una API o servicio web,
	// establecer el encabezado Content-Type a application/json asegura que
	// cualquier cliente que consuma la API entienda que el formato de la respuesta será JSON.
	w.Header().Set("Content-Type", "application/json")

	// Este método se utiliza para convertir un objeto de Go a JSON y enviarlo como respuesta HTTP.
	json.NewEncoder(w).Encode(movies)
	/*
		Creación del Codificador (Encoder):
		json.NewEncoder(w): Crea un nuevo Encoder que escribe la salida JSON en w,
		donde w es un http.ResponseWriter. http.ResponseWriter se utiliza para enviar
		la respuesta de vuelta al cliente.

		Codificación y Envío del Objeto:
		.Encode(movies): Convierte el objeto movies (una lista de películas en tu caso) a JSON y
		lo escribe en la respuesta HTTP.
	*/
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		Cuando defines rutas en Gorilla Mux con variables de ruta, como /{id} en la URL /movies/{id},
		mux puede capturar esas variables y permitirte acceder a ellas fácilmente
		dentro de tus manejadores de solicitudes. La función mux.Vars(r) extrae estas
		variables de la solicitud entrante y las devuelve como un mapa (map[string]string).
	*/
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
	params := mux.Vars(r)

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
			// Este método se utiliza para convertir datos JSON recibidos en una solicitud HTTP en un objeto de Go.
			_ = json.NewDecoder(r.Body).Decode(&movie)
			/*
				Creación del Decodificador (Decoder):
				json.NewDecoder(r.Body): Crea un nuevo Decoder que lee la entrada JSON de r.Body,
				donde r.Body es el cuerpo de la solicitud HTTP.
				Este cuerpo de la solicitud contiene los datos JSON enviados por el cliente.

				Decodificación del JSON en un Objeto:
				.Decode(&movie): Convierte los datos JSON en el cuerpo
				de la solicitud al objeto movie (una instancia de la estructura Movie)
			*/
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
