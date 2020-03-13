package book

type Book struct {
	Id        int    `json:"id"`
	BookName  string `json:"book_name"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}

var BookIDCount = 0
var Books = make(map[int]Book)

func (book *Book) addBook() {

	BookIDCount++
	book.Id = BookIDCount
	book.Available = true

	Books[BookIDCount] = Book{
		Id:        book.Id,
		BookName:  book.BookName,
		Author:    book.Author,
		Available: book.Available,
	}

}
