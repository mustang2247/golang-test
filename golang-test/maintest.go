package main

import "fmt"
import (
	"./src/db/mysql"
)

func main()  {
	var a int = 65
	// 表示65的文本A
	b := string(a)
	fmt.Println(b)

	//db, err := sql.Open("mysql", "root:mustang@tcp(localhost:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci")
	err := mysql.Open("root", "mustang", "localhost", "3306", "test")
	if err != nil {
		fmt.Errorf(err.Error() + "\n")
	}

	var sqlcmd string = "INSERT INTO user(username, password) VALUES(?, ?)"
	mysql.Insert(sqlcmd, "guotie", "guotie")

}