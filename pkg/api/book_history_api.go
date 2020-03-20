package api

import (
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/Emon331046/libraryManagement/pkg/db"
	"github.com/Emon331046/libraryManagement/pkg/model"
)

func AddNewPurchase(ctx *macaron.Context, bookHistory model.BookHistoryDb) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := ctx.Req.Header.Get("current_user_type")
	if currentUserType != "admin" {
		ctx.JSON(http.StatusBadGateway, "only admin can update purchase book info")
		return
	}
	userId := bookHistory.UserId
	bookId := bookHistory.BookId
	resultPurchase, err := db.AddNewPurchase(userId, bookId)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resultPurchase)

}

func ReturnBook(ctx *macaron.Context, bookHistory model.BookHistoryDb) {
	//_, err := strconv.Atoi(r.Header.Get("current_user_id"))
	currentUserType := ctx.Req.Header.Get("current_user_type")
	if currentUserType != "admin" {

		ctx.JSON(http.StatusUnauthorized, "user type didn't match")
		return
	}

	userId := bookHistory.UserId
	bookId := bookHistory.BookId

	result, err := db.ReturnBookMethod(userId, bookId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, "can't execute return request")
	}
	ctx.JSON(http.StatusCreated, result)

}
