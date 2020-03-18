package db

import (
	"errors"
	"fmt"

	"github.com/emon331046/libraryManagement/pkg/model"
)

func AddNewPurchase(userId int, bookId int) (*model.BookHistoryDb, error) {
	var user model.UserDb
	var book model.Bookdb
	eng.ID(userId).Get(&user)
	eng.ID(bookId).Get(&book)
	if book.NotAvailable {
		return nil, errors.New("book not available")
	}
	bookHistory := model.BookHistoryDb{
		UserId:   user.ID,
		UserName: user.Name,
		UserMail: user.Mail,
		BookId:   book.Id,
		BookName: book.BookName,
	}
	book.NotAvailable = true
	//fmt.Println(book)

	session := eng.NewSession()
	defer session.Close()
	err := session.Begin()

	_, err1 := eng.Id(bookId).UseBool().Update(&book)

	if err1 != nil {
		session.Rollback()
		return nil, err1
	}

	_, err2 := eng.Insert(bookHistory)

	if err2 != nil {
		session.Rollback()
		return nil, err
	}
	err3 := session.Commit()

	if err3 != nil {
		return nil, errors.New("server failed")
	}
	eng.Where("user_id =? AND book_id = ? AND returned= FALSE", userId, bookId).Get(&bookHistory)

	return &bookHistory, nil

}

func ReturnBookMethod(userId int, bookId int) (*model.BookHistoryDb, error) {

	var bookHistory model.BookHistoryDb
	okk, err0 := eng.Where("book_id=? AND user_id=? AND returned = FALSE", bookId, userId).Get(&bookHistory)
	if err0 != nil {
		return nil, err0

	}
	if okk {

		var book model.Bookdb
		eng.Id(bookId).Get(&book)
		fmt.Println(book)
		book.NotAvailable = false
		_, err := eng.Id(bookId).UseBool().Update(book)
		fmt.Println(book)

		_, err1 := eng.Where("book_id=? AND user_id=?", bookId, userId).Get(&bookHistory)

		bookHistory.Returned = true
		_, err2 := eng.Id(bookHistory.HistoryId).UseBool().Update(&bookHistory)

		if err2 != nil {
			return nil, err1
		}

		_, err3 := eng.Where("book_id=? AND user_id=?", bookId, userId).UseBool().Update(&bookHistory)

		if err3 != nil {
			return nil, err
		}

		eng.Where("user_id =? AND book_id = ?", userId, bookId).Get(&bookHistory)

		return &bookHistory, nil

	}
	return nil, errors.New("no returned book data found")
}
