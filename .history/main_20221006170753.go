package main

import(
"fmt"
"encoding/json"
"net/http"
"math/rand"
"strconv"
"github.com/gorilla/mux"
)

func main(){
	// init router
	r :=mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books/". getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}". getBook).Methods("GET")
	r.HandleFunc("/api/books/". createBook).Methods("GET")
	r.HandleFunc("/api/books/". getBooks).Methods("GET")
	r.HandleFunc("/api/books/". getBooks).Methods("GET")
}