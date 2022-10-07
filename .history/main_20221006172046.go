package main

import(
"fmt"
"encoding/json"
"net/http"
"math/rand"
"strconv"
"github.com/gorilla/mux"
)

// Book struct (Model) class
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request){

}
func getBook(w http.ResponseWriter, r *http.Request){

}
func createBook(w http.ResponseWriter, r *http.Request){

}
func updateBook(w http.ResponseWriter, r *http.Request){

}
func deleteBook(w http.ResponseWriter, r *http.Request){

}


func main(){
	// init router
	r :=mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}