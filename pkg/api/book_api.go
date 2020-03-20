package api

import (
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/macaron.v1"

	"github.com/Emon331046/libraryManagement/pkg/db"
	"github.com/Emon331046/libraryManagement/pkg/model"
)

func AddNewBook(ctx *macaron.Context, book model.Bookdb) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := ctx.Req.Header.Get("current_user_type")
	if currentUserType != "admin" {

		ctx.JSON(http.StatusBadGateway, "unmatched type ")
		return
	}

	fmt.Println(book)

	resultBook, err := db.AddBook(book)
	if err != nil {

		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, resultBook)

}
func ShowAllBooks(ctx *macaron.Context) {

	resultBooks, err := db.ShowAllBooks()
	if err != nil {

		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultBooks)
	return

}
func ShowBook(ctx *macaron.Context) {
	key := ctx.Params(":bookId")
	fmt.Println(key)
	bookId, err1 := strconv.Atoi(key)
	fmt.Println("*********hello from show book**********", key, bookId)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, err1.Error())
		return
	}
	resultBook, err := db.ShowBook(bookId)
	if err != nil {

		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultBook)

}

func DeleteBook(ctx *macaron.Context) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := ctx.Req.Header.Get("current_user_type")
	if currentUserType != "admin" {

		ctx.JSON(http.StatusBadRequest, "user is not admin")
		return
	}

	//book_id, _ := strconv.Atoi(mux.Vars(r)["book_id"])
	book_id, err2 := strconv.Atoi(ctx.Params(":bookId"))
	if err2 != nil {

		ctx.JSON(http.StatusBadGateway, err2.Error())
	}
	fmt.Println(book_id)

	_, err := db.DeleteBookMethod(book_id)
	if err == nil {
		ctx.JSON(http.StatusResetContent, "book has been deleted")

		return
	}
	ctx.JSON(http.StatusBadGateway, err.Error())
	return

}
