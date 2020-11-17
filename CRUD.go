package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Book struct {
    BookID  string  `json:"id"`
    Title   string  `json:"title"`
    Desc    string  `json:"description"`
    Content string  `json:"content"`
}

var Books []Book

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllBooks(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllBooks")
    json.NewEncoder(w).Encode(Books)
}

func returnBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["bookId"]
    fmt.Println("BookID : ", key)
    for _, book := range Books {
        if book.BookID == key {
            json.NewEncoder(w).Encode(book)
        }
    }
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &book)
    Books = append(Books, book)

    json.NewEncoder(w).Encode(Books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["bookId"]
    for index, book := range Books {
        if book.BookID == id {
            Books = append(Books[:index], Books[index+1:]...)
        }
    }

    json.NewEncoder(w).Encode(Books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    var newbook Book
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &newbook)
    for index, book := range Books {
        if book.BookID == newbook.BookID {
            Books = append(Books[:index], Books[index+1:]...)
        }
    }
    Books = append(Books, newbook)

    json.NewEncoder(w).Encode(Books)
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/allbooks", returnAllBooks)
    myRouter.HandleFunc("/book/{bookId}", returnBookById)
    myRouter.HandleFunc("/newbook", createNewBook).Methods("POST")
    myRouter.HandleFunc("/delbook/{bookId}", deleteBook).Methods("DELETE")
    myRouter.HandleFunc("/updatebook", updateBook).Methods("PUT")
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Books = []Book{
        Book{BookID: "1", Title: "Machine Learning Basic", Desc: "Book Description", Content: "Book Content"},
        Book{BookID: "2", Title: "What is Data Science?", Desc: "Book Description", Content: "Book Content"},
    }
    handleRequests()
}
