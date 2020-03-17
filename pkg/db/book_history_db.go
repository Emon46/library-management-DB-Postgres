package db

type BookHistoryDb struct {
	HistoryId int    `xorm:"pk autoincr history_id1"`
	BookId    int    `xorm:"book_id"`
	BookName  string `xorm:"book_name'"`
	UserId    int    `xorm:"user_id"`
	UserName  string `xorm:"user_name_hello"`
	Returned  bool   `xorm:"returned DEFAULT'true'"`

	PurchasedDate string `xorm:"created"`
	ReturnDate    string `xorm:"updasate updated "`
}

func (BookHistoryDb) TableName() string {
	//fmt.Println(BookHistoryDb{})
	return "book_history"
}
