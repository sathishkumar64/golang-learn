package main
import (
	"log"
	"net/http"	
	"github.com/gorilla/mux"
	"encoding/json"
)

type Book struct{
	ID string  `json:"id"`
	Isbn string `json:"isbn"`
	Title string	`json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

var books []Book

func main(){
 
	// init router
	r := mux.NewRouter() 

	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})


	//Router handler ,endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
//	r.HandleFunc("api/book/{id}",getBook).Methods("GET")
//	r.HandleFunc("api/books",createBooks).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}