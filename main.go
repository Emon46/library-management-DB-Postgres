package main

import (
	"github.com/gorilla/mux"
	"libraryManagement/pkg"
	"log"
	"net/http"
	"time"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", pkg.Login)
	r.HandleFunc("/register", pkg.Register)
	http.Handle("/",r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}


	log.Fatal(srv.ListenAndServe())

}