package main

import(
// "fmt"
"encoding/json"
"log"
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

// Init books var as a slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r) //GETs PARAMs
	//Loop through books and validate id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID =strconv.Itoa(rand.Intn(1000000)) //mock ID not safe
	books = append(books,book)
	json.NewEncoder(w).Encode(book)


}
func updateBook(w http.ResponseWriter, r *http.Request){

}
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books)
		}

	}
}


func main(){
	// init router
	r :=mux.NewRouter()

	// Moc data @TOdo impleemtn DB
	books = append(books, Book{ID: "1", Isbn:"33452", Title:"Book One", Author:
	 &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn:"34452", Title:"Book Two", Author: 
	&Author{Firstname: "Ken", Lastname: "Jon"}})
	books = append(books, Book{ID: "3", Isbn:"35452", Title:"Book Three", Author: 
	&Author{Firstname: "Paul", Lastname: "Walker"}})
	books = append(books, Book{ID: "4", Isbn:"36452", Title:"Book Four", Author: 
	&Author{Firstname: "King", Lastname: "Kong"}})
	// Route Handlers / Endpoints
	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}