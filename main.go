package main

import (
	"log"
	"net/http"
	"time"

	"github.com/emon331046/libraryManagement/pkg/middleware"

	"github.com/emon331046/libraryManagement/pkg/api"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.Use(middleware.JwtMiddleWare)

	r.HandleFunc("/login", api.Login).
		Methods("GET")
	r.HandleFunc("/register", api.Register).
		Methods("POST")
	r.HandleFunc("/user-profile/{user_id}", api.UserProfile).
		Methods("Get")
	r.HandleFunc("/edit-profile", api.EditUserProfile).
		Methods("PATCH")
	r.HandleFunc("/book", api.AddNewBook).
		Methods("POST")
	r.HandleFunc("/book", api.ShowAllBooks).
		Methods("Get")
	r.HandleFunc("/book/{book_id}", api.ShowBook).
		Methods("Get")
	r.HandleFunc("/purchase-book", api.AddNewPurchase).
		Methods("POST")
	r.HandleFunc("/return-book", api.ReturnBook).
		Methods("PUT")
	r.HandleFunc("/delete-book/{book_id}", api.DeleteBook)
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
