package main

import (
	"libraryManagement/pkg"
	"libraryManagement/pkg/book"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", pkg.Login).
		Methods("GET")
	r.HandleFunc("/register", pkg.Register).
		Methods("POST")
	r.HandleFunc("/book", book.AddNewBook).
		Methods("POST")
	r.HandleFunc("/book", book.ShowAllBooks).
		Methods("Get")
	r.HandleFunc("/purchase-book", book.AddNewPurchase).
		Methods("POST")
	r.HandleFunc("/return-book", book.ReturnBook).
		Methods("PUT")
	r.HandleFunc("/delete-book/{book_id}", book.DeleteBook)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
