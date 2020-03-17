package db

import "fmt"

type BookDb struct {
	Id        int    `xorm:"id"`
	BookName  string `xorm:"book_name"`
	Author    string `xorm:"author"`
	Available bool   `xorm:"available"`
}

func (BookDb) TableName() string {
	fmt.Println(BookDb{})
	return "books"
}
