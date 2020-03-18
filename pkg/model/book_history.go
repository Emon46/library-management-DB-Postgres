package model

type BookHistoryDb struct {
	HistoryId int    `xorm:"pk autoincr history_id" json:"history_id"`
	BookId    int    `xorm:"book_id" json:"book_id"`
	BookName  string `xorm:"book_name" json:"book_name"`
	UserId    int    `xorm:"user_id" json:"user_id"`
	UserName  string `xorm:"user_name" json:"user_name"`
	UserMail  string `xorm:"user_mail" json:"user_mail"`
	Returned  bool   `xorm:"returned DEFAULT FALSE" json:"returned"`

	PurchasedDate string `xorm:"created" json:"purchased_date"`
	ReturnDate    string `xorm:"update updated " json:"return_date"`
}

func (BookHistoryDb) TableName() string {
	//fmt.Println(BookHistoryDb{})
	return "book_history"
}
