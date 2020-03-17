package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emon331046/libraryManagement/pkg"
	"github.com/gorilla/mux"
)

func ShowAllBooks(w http.ResponseWriter, r *http.Request) {
	myResponse := pkg.MyData{
		Status:  http.StatusOK,
		Error:   nil,
		Message: "created new user",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myResponse)

}

func ShowBook(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["book_id"]
	bookId, err := strconv.Atoi(key)
	if err != nil {
		myResponse := pkg.MyData{
			Status:  http.StatusBadRequest,
			Error:   nil,
			Message: "no match for this book",
			Success: "true",
			Data:    Books[bookId],
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(myResponse)
	}
	fmt.Println(bookId, "dfj", Books[bookId].Id)

	if Books[bookId].Id == bookId {
		//fmt.Println("hello9")

		myResponse := pkg.MyData{
			Status:  http.StatusOK,
			Error:   nil,
			Message: "created new user",
			Success: "true",
			Data:    Books[bookId],
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(myResponse)
		//return
	} else {
		//fmt.Println("error9")

		//fmt.Println(bookId, Books[bookId].Id)
		myResponse := pkg.MyData{
			Status:  http.StatusBadRequest,
			Error:   err,
			Message: "No Book found",
			Success: "false",
			Data:    "{}",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(myResponse)
	}
	//fmt.Println("sdgh")

}
func AddNewBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {
		myResponse := pkg.MyData{
			Status:  http.StatusNotAcceptable,
			Error:   errors.New("user type didn't match"),
			Message: "only admin can add new book ",
			Success: "true",
			Data:    Books,
		}

		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(myResponse)
		return
	}
	var book Book
	err1 := json.NewDecoder(r.Body).Decode(&book)
	if err1 != nil {
		return
	}
	for _, bookVar := range Books {
		if bookVar.BookName == book.BookName && bookVar.Author == book.Author {
			return
		}
	}

	book.addBook()

	myResponse := pkg.MyData{
		Status:  http.StatusCreated,
		Error:   nil,
		Message: "created new user",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
func AddNewPurchase(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {
		myResponse := pkg.MyData{
			Status:  http.StatusNotAcceptable,
			Error:   errors.New("user type didn't match"),
			Message: "only admin can update purchase book info",
			Success: "true",
			Data:    Books,
		}

		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(myResponse)
		return
	}
	var bookHistory BookHistory
	json.NewDecoder(r.Body).Decode(&bookHistory)
	bookHistory.NewPurchase()

	myResponse := pkg.MyData{
		Status:  http.StatusCreated,
		Error:   nil,
		Message: "added this book to purchase list",
		Success: "true",
		Data:    BooksHistoryList,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
func ReturnBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {
		myResponse := pkg.MyData{
			Status:  http.StatusNotAcceptable,
			Error:   errors.New("user type didn't match"),
			Message: "only admin can update return book info",
			Success: "true",
			Data:    Books,
		}

		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(myResponse)
		return
	}
	//fmt.Println("kdsfh")
	var bookHistory BookHistory
	json.NewDecoder(r.Body).Decode(&bookHistory)
	bookHistory.ReturnBookMethod()

	myResponse := pkg.MyData{
		Status:  http.StatusCreated,
		Error:   nil,
		Message: "added this book to purchase list",
		Success: "true",
		Data:    BooksHistoryList,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := r.Header.Get("current_user_type")
	if currentUserType != "admin" {
		myResponse := pkg.MyData{
			Status:  http.StatusResetContent,
			Error:   errors.New("user type didn't match"),
			Message: "only admin can delete",
			Success: "true",
			Data:    Books,
		}

		w.WriteHeader(http.StatusResetContent)
		json.NewEncoder(w).Encode(myResponse)
		return
	}

	book_id, _ := strconv.Atoi(mux.Vars(r)["book_id"])
	fmt.Println(book_id)

	delete(Books, book_id)

	delete(BooksHistoryList, book_id)
	//fmt.Println(BooksHistoryList)

	myResponse := pkg.MyData{
		Status:  http.StatusResetContent,
		Error:   nil,
		Message: "successfully deleted",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusResetContent)
	json.NewEncoder(w).Encode(myResponse)
	//another approach
	//data, err := json.Marshal(myResponse)
	//w.Write(data)

}
