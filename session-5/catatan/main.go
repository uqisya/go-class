package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Book struct {
	ID     int
	Title  string
	Stock  int
	Author string
}

var books = []Book{
	{ID: 1, Title: "Rich Dad Poor Dad", Stock: 20, Author: "Test Author"},
	{ID: 2, Title: "Make It Happen Now", Stock: 15, Author: "Prita Ghozie"},
}

var PORT = ":8080"

func main() {
	// routing
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/book", createBook)

	fmt.Println("App is listening on port ", PORT)
	http.ListenAndServe(PORT, nil)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	// method header -> bentuk dari data response yg nantinya kirim ke client
	// atur content type nya jadi application/json
	// w.Header().Set("Content-Type", "application/json")

	// hit API nya dari sini
	if r.Method == "GET" {
		// get books
		// json.NewEncoder(w).Encode(books)
		// return

		// kalo pake parsing file
		tpl, err := template.ParseFiles("html/template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// execute file html, dan kirim data apa
		tpl.Execute(w, books)
		return
	}

	// if error
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		// create book
		title := r.FormValue("title")
		stock := r.FormValue("stock") // must convert str to int
		author := r.FormValue("author")

		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid Stock", http.StatusBadRequest)
		}

		newBook := Book{
			ID:     len(books) + 1,
			Title:  title,
			Stock:  convertStock,
			Author: author,
		}

		books = append(books, newBook)

		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
