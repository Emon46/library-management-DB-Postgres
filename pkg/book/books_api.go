package book

import (
	"encoding/json"
	"fmt"
	"libraryManagement/pkg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ShowAllBooks(w http.ResponseWriter, r *http.Request) {
	myResponse := pkg.MyData{
		Status:  http.StatusOK,
		Error:   "null",
		Message: "created new user",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myResponse)

}
func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
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
		Error:   "null",
		Message: "created new user",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)
}
func AddNewPurchase(w http.ResponseWriter, r *http.Request) {
	var bookHistory BookHistory
	json.NewDecoder(r.Body).Decode(&bookHistory)
	bookHistory.NewPurchase()

	myResponse := pkg.MyData{
		Status:  http.StatusCreated,
		Error:   "null",
		Message: "added this book to purchase list",
		Success: "true",
		Data:    BooksHistoryList,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
func ReturnBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kdsfh")
	var bookHistory BookHistory
	json.NewDecoder(r.Body).Decode(&bookHistory)
	bookHistory.ReturnBookMethod()

	myResponse := pkg.MyData{
		Status:  http.StatusCreated,
		Error:   "null",
		Message: "added this book to purchase list",
		Success: "true",
		Data:    BooksHistoryList,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myResponse)

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	book_id, _ := strconv.Atoi(mux.Vars(r)["book_id"])
	fmt.Println(book_id)

	delete(Books, book_id)

	delete(BooksHistoryList, book_id)
	//fmt.Println(BooksHistoryList)

	myResponse := pkg.MyData{
		Status:  http.StatusResetContent,
		Error:   "null",
		Message: "successfully deleted",
		Success: "true",
		Data:    Books,
	}

	w.WriteHeader(http.StatusResetContent)
	json.NewEncoder(w).Encode(myResponse)
	//data, err := json.Marshal(myResponse)
	//w.Write(data)

}
