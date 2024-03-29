package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Structs-model
type Book struct{
	ID		string 		`json:id`	
	Isbn		string 		`json:isbn`	
	Title		string 		`json:title`	
	Author		*Author 		`json:author`	
}

type Author struct{
	Firstname  	string 		`json:firstname`
	Lastname  	string 		`json:lastname`
}

//Init books var as a slice book struct
var books  []Book

//get all books
func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//get a book
func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) //get params

	//start loop through books
	for _,item:= range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create a book
func createBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

//Update a book
func updateBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index,item := range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete a book
func deleteBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index,item := range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000))
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main(){
	//Init Router
	r:= mux.NewRouter()

	//mock data - @todo implement db
	books= append(books, Book{ID: "1", Isbn: "23423", Title:"book 1", Author: &Author{Firstname: "John", Lastname: "Doe" }})
	books= append(books, Book{ID: "2", Isbn: "23424", Title:"book 2", Author: &Author{Firstname: "Jahn", Lastname: "Dae" }})
	books= append(books, Book{ID: "3", Isbn: "23425", Title:"book 3", Author: &Author{Firstname: "Johnny", Lastname: "Doc"}})
	books= append(books, Book{ID: "4", Isbn: "23426", Title:"book 4", Author: &Author{Firstname: "Join", Lastname: "Dos" }})
	//route handler
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))

}
