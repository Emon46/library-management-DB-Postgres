package db

import (
	"errors"
	"fmt"

	"github.com/Emon331046/libraryManagement/pkg/model"
)

func AddBook(book model.Bookdb) (*model.Bookdb, error) {

	_, err := eng.Insert(book)
	if err != nil {
		return nil, err

	}
	eng.Where("book_name=?", book.BookName).Get(&book)

	return &book, nil

}
func ShowAllBooks() (*model.Booksdb, error) {
	booksdb := model.Booksdb{}
	var books []model.Bookdb
	eng.Find(&books)
	fmt.Println(books)
	booksdb.Books = append(booksdb.Books, books...)
	fmt.Println(booksdb)
	return &booksdb, nil
}
func ShowBook(bookId int) (*model.Bookdb, error) {
	var book model.Bookdb
	okk, _ := eng.Where("id=?", bookId).Get(&book)
	//println(okk, err)
	if okk {
		return &book, nil
	}
	return nil, errors.New("the book didn't find")
}
func DeleteBookMethod(bookId int) (bool, error) {
	session := eng.NewSession()
	defer session.Close()
	err := session.Begin()

	okk, err := eng.Id(bookId).Delete(&model.Bookdb{})
	if err != nil {
		session.Rollback()
		return false, errors.New("roll backed")
	}
	_, err1 := eng.Where("book_id =?", bookId).Delete(&model.BookHistoryDb{})

	if err1 != nil {
		session.Rollback()
		return false, errors.New("roll backed")
	}
	err2 := session.Commit()
	if err2 != nil {
		return false, errors.New("server failed")
	}
	if okk > 0 {
		return true, nil
	}

	return false, errors.New("no book found last one")
}
