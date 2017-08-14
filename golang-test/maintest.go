package main

import (
	"./src/db/mysql"
	"fmt"
)

func main()  {
	//db, err := sql.Open("mysql", "root:mustang@tcp(localhost:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci")
	err := mysql.Open("root", "mustang", "localhost", "3306", "test")
	if err != nil {
		fmt.Errorf(err.Error() + "\n")
	}

	var sqlcmd string = "INSERT INTO user(username, password) VALUES(?, ?)"
	mysql.Insert(sqlcmd, "guotie", "guotie")

	res := mysql.Query("SELECT * FROM user", nil)
	//fmt.Println(res)
	if res != nil {
		for k, v := range res{
			fmt.Println(k, v)
		}
	}

}