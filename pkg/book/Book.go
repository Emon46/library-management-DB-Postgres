package book

type Book struct {
	Id        int    `json:"id"`
	BookName  string `json:"book_name"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}

var bookIDCount = 0
var Books = make(map[int]Book)

func (book *Book) addBook() {

	bookIDCount++
	book.Id = bookIDCount
	book.Available = true

	Books[bookIDCount] = Book{
		Id:        book.Id,
		BookName:  book.BookName,
		Author:    book.Author,
		Available: book.Available,
	}

}
