package db

import (
	"fmt"

	"github.com/Emon331046/libraryManagement/pkg/model"

	"xorm.io/core"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var (
	eng    *xorm.Engine
	tables []interface{}
)
var DBUserName = "postgres"
var DBPort = 5432
var DBName = "library_management"

func init() {
	tables = append(tables, new(model.UserDb),
		new(model.Bookdb),
		new(model.BookHistoryDb))

	var err error
	eng, err = GetPostgresClient()
	if err != nil {
		fmt.Println("ORM is not connected\n Error:", err)
	}

	eng.SetTableMapper(core.SameMapper{})
	eng.SetColumnMapper(core.SnakeMapper{})
	eng.Sync2(tables...)
	//eng.Table("usssser").CreateTable(new(model.UserDb))

}
func GetPostgresClient() (*xorm.Engine, error) {
	fmt.Println(".....ok......")
	cnnstr := fmt.Sprintf("user=%s host=127.0.0.1 port=%v dbname=%s sslmode=disable", DBUserName, DBPort, DBName)
	return xorm.NewEngine("postgres", cnnstr)
}
