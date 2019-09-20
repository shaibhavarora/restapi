package main
 import (
	 "encoding/json"
	 _ "encoding/json"
	 _ "fmt"
	 _ "fmt"
	 "github.com/gorilla/mux"
	 _ "github.com/gorilla/mux"
	 "log"
	 "math/rand"
	 _ "math/rand"
	 "net/http"
	 _ "net/http"
	 "strconv"
	 _ "strconv"
 )

//Books structs (model)
 type Book struct {
 	 ID string `json:"id"`
	 ISBN string `json:"isbn"`
	 Title string `json:"title"`
	 Author *Author `json:"author"`
 }


//Author struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Init books var as a slice books struct

var books []Book



//get all books

//route handler function

func getBooks(w http.ResponseWriter, r *http.Request){
	//setting header value of json else it is gonna be served as text
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)

}
//get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) //get the params
	//loop through the books and find the correct id

	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ =json.NewDecoder(r.Body).Decode(&book)
	book.ID =strconv.Itoa(rand.Intn(100000)) //rand intn should be an string and that's what we are doing with it//mock id:unsafe
	books =append(books, book)
	json.NewEncoder(w).Encode(book)

}
//update the book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ =json.NewDecoder(r.Body).Decode(&book)
			book.ID =params["id"]
			books =append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
//delete the book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main(){
	//init router

	r := mux.NewRouter()

	//Mock data
	books =append(books, Book{ID:"1",ISBN:"55678",Title:"book1", Author: &Author{Firstname:"John",Lastname:"Doe"}})
	books =append(books, Book{ID:"2",ISBN:"55578",Title:"book2", Author: &Author{Firstname:"John",Lastname:"Doe2"}})
	books =append(books, Book{ID:"3",ISBN:"55120",Title:"book3", Author: &Author{Firstname:"John",Lastname:"Doe3"}})
	books =append(books, Book{ID:"4",ISBN:"55000",Title:"book4", Author: &Author{Firstname:"John",Lastname:"Doe4"}})

	//Route handler
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	//r.HandleFunc("/api/exceptions/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000",r))
	log.Fatal(http.ListenAndServe(":8000",r))


}