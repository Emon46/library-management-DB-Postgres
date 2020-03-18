package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emon331046/libraryManagement/pkg/db"
	"github.com/emon331046/libraryManagement/pkg/model"
)

func AddNewPurchase(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {

		http.Error(w, "only admin can update purchase book info", http.StatusBadGateway)
		return
	}
	var bookHistory model.BookHistoryDb
	json.NewDecoder(r.Body).Decode(&bookHistory)
	userId := bookHistory.UserId
	bookId := bookHistory.BookId
	resultPurchase, err := db.AddNewPurchase(userId, bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultPurchase)

}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {

		http.Error(w, "user type didn't match", http.StatusUnauthorized)
		return
	}
	//fmt.Println("kdsfh")

	var bookHistory model.BookHistoryDb
	json.NewDecoder(r.Body).Decode(&bookHistory)
	userId := bookHistory.UserId
	bookId := bookHistory.BookId
	fmt.Println(bookId)

	result, err := db.ReturnBookMethod(userId, bookId)
	if err != nil {
		http.Error(w, "can't execute return request", http.StatusBadGateway)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
