package book

import "fmt"

type BookHistory struct {
	HistoryId     int    `json:"history_id"`
	BookId        int    `json:"book_id"`
	BookName      string `json:"book_name"`
	UserId        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	PurchasedDate string `json:"purchased_date"`
	ReturnDate    string `json:"return_date"`
	Returned      bool   `json:"returned"`
}

var HistoryCount = 0
var BooksHistoryList = make(map[int][]BookHistory)

func (bookHistory *BookHistory) NewPurchase() {

	HistoryCount++
	bookHistory.HistoryId = HistoryCount
	bookHistory.ReturnDate = "doesn't returned yet"

	BooksHistoryList[bookHistory.BookId] = append(BooksHistoryList[bookHistory.BookId],
		BookHistory{
			HistoryId:     bookHistory.HistoryId,
			BookId:        bookHistory.BookId,
			BookName:      bookHistory.BookName,
			UserId:        bookHistory.UserId,
			UserName:      bookHistory.UserName,
			PurchasedDate: bookHistory.PurchasedDate,
			ReturnDate:    bookHistory.ReturnDate,
			Returned:      false,
		})
	for i, bookVar := range Books {
		if bookVar.Id == bookHistory.BookId {
			fmt.Println("fg")
			bookVar.Available = false
			Books[i] = bookVar

		}
	}

}

func (bookHistory *BookHistory) ReturnBookMethod() {

	fmt.Println("sdkdsv")
	for i, bookHistoryVar := range BooksHistoryList[bookHistory.BookId] {
		fmt.Println(bookHistory.UserId, bookHistoryVar.UserId)
		if bookHistoryVar.UserId == bookHistory.UserId && bookHistoryVar.Returned == false {
			fmt.Println("fghg")
			bookHistoryVar.Returned = true
			bookHistoryVar.ReturnDate = bookHistory.ReturnDate
			BooksHistoryList[bookHistory.BookId][i] = bookHistoryVar
			for _, bookVar := range Books {
				if bookVar.Id == bookHistory.BookId {
					fmt.Println("fg")
					bookVar.Available = true

					Books[bookVar.Id] = bookVar

				}
			}
		}
	}

}
