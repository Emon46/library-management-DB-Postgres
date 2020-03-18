package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/emon331046/libraryManagement/pkg/db"
	"github.com/emon331046/libraryManagement/pkg/model"
)

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {

		http.Error(w, "unmatched type ", http.StatusBadGateway)
		return
	}
	var book model.Bookdb
	err1 := json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	if err1 != nil {
		return
	}

	resultBook, err := db.AddBook(book)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultBook)

}
func ShowAllBooks(w http.ResponseWriter, r *http.Request) {

	resultBooks, err := db.ShowAllBooks()
	if err != nil {

		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(resultBooks)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultBooks)

}
func ShowBook(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["book_id"]
	bookId, err1 := strconv.Atoi(key)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadGateway)
		return
	}
	resultBook, err := db.ShowBook(bookId)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultBook)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {

		http.Error(w, "user is not admin", http.StatusBadRequest)
		return
	}

	book_id, _ := strconv.Atoi(mux.Vars(r)["book_id"])
	fmt.Println(book_id)

	_, err := db.DeleteBookMethod(book_id)
	if err == nil {
		w.WriteHeader(http.StatusResetContent)

		w.Write([]byte("book has been deleted"))
		return
	}
	http.Error(w, err.Error(), http.StatusBadGateway)
	return

}
