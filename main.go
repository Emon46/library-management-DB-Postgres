package main

import (
	"libraryManagement/pkg"
	"libraryManagement/pkg/book"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Middleware(next http.Handler) http.Handler {
	println("hello")
	return next
}
func init() {
	pkg.Users = append(pkg.Users, pkg.User{
		ID:       pkg.UserIdCount,
		Name:     "demo user",
		Mail:     "demouser@mail.com",
		Password: "123456",
		PhoneNo:  "019999999",
		UserType: "user",
	})
	pkg.UserIdCount++
	pkg.Users = append(pkg.Users, pkg.User{
		ID:       pkg.UserIdCount,
		Name:     "admin",
		Mail:     "admin@mail.com",
		Password: "123456",
		PhoneNo:  "019999999",
		UserType: "admin",
	})
	pkg.UserIdCount++

}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", pkg.Login).
		Methods("GET")
	r.HandleFunc("/register", pkg.Register).
		Methods("POST")
	r.HandleFunc("/user/{user_id}", pkg.UserProfile).
		Methods("Get")
	r.Handle("/book", AuthMiddleware(book.AddNewBook(), "admin")).
		Methods("POST")
	r.HandleFunc("/book", book.ShowAllBooks).
		Methods("Get")
	r.HandleFunc("/book/{book_id}", book.ShowBook).
		Methods("Get")
	//http.HandleFunc("/df",fu)

	r.Handle("/purchase-book", AuthMiddleware(book.AddNewPurchase(), "admin")).
		Methods("POST")
	r.Handle("/return-book", AuthMiddleware(book.ReturnBook(), "admin")).
		Methods("PUT")
	r.Handle("/delete-book/{book_id}", AuthMiddleware(book.DeleteBook(), "admin"))
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
