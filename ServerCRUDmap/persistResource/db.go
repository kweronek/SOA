package persistResource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDatabase() {
	fmt.Println("try to connect database!")
	db, err = sql.Open(
		"mysql",
		"root:root@/test",
	)
	if err != nil {
		panic(err.Error())
	} else {
		defer db.Close()
		fmt.Println("Database connected! (...not really...)")
	}

}
