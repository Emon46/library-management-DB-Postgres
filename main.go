package main

import (
	"log"
	"net/http"
	"time"

	"github.com/emon331046/libraryManagement/pkg/api"

	"github.com/emon331046/libraryManagement/pkg"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.Use(pkg.JwtMiddleWare)

	//r.HandleFunc("/login", pkg.Login).
	//	Methods("GET")
	r.HandleFunc("/register", api.Register).
		Methods("POST")
	r.HandleFunc("/user-profile/{user_id}", api.UserProfile).
		Methods("Get")
	//r.HandleFunc("/edit-profile/{user_id}", pkg.EditUserProfile).
	//	Methods("PATCH")
	////r.Handle("/edit-profile/{user_id}", AuthMiddleware(pkg.EditUserProfile(), "user")).
	////	Methods("PATCH")
	//r.HandleFunc("/book", book.AddNewBook).
	//	Methods("POST")
	//r.HandleFunc("/book", book.ShowAllBooks).
	//	Methods("Get")
	//r.HandleFunc("/book/{book_id}", book.ShowBook).
	//	Methods("Get")
	////http.HandleFunc("/df",fu)
	//
	//r.HandleFunc("/purchase-book", book.AddNewPurchase).
	//	Methods("POST")
	//r.HandleFunc("/return-book", book.ReturnBook).
	//	Methods("PUT")
	//r.HandleFunc("/delete-book/{book_id}", book.DeleteBook)
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
