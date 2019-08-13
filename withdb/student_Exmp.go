package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"github.com/gorilla/mux"
	"context"
	"fmt"
	
)

/*Book Object */
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
var client *mongo.Client

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	collection := client.Database("golearn").Collection("book")
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		cursor.Decode(&book)
		books = append(book, book)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(books)
}

func main(){
	fmt.Println("Applicaiton is starting...")
	ctx,err := context.WithTimeout(context.Background(),10*time.Second)
	client,err := mongo.Connect(ctx,"mongodb://localhost:27017")
	r := mux.NewRouter() 
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	http.ListenAndServe(":8080", r)
}