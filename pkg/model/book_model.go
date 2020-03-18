package model

import "time"

type Bookdb struct {
	Id           int       `xorm:"id pk autoincr" json:"id"`
	BookName     string    `xorm:"book_name unique" json:"book_name"`
	Author       string    `xorm:"author" json:"author"`
	NotAvailable bool      `xorm:"not_available DEFAULT FALSE" json:"not_available"`
	CreatedAt    time.Time `xorm:"created" json:"created_at" `
}

type Booksdb struct {
	Books []Bookdb
}

func (Booksdb) TableName() string {
	return "books"
}
func (Bookdb) TableName() string {
	return "books"
}
