package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	stmt.Exec("guotie", "guotie")
	stmt.Exec("testuser", "123123")
}

func selectData (db *sql.DB, uname string)  {
	//rows, err := db.Query("select username from user where username = ?", uname)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer rows.Close()
	//
	//var boo = false
	//if(rows.Next()){
	//	boo = true
	//}

}

func main() {
	//数据库连接
	db, err := sql.Open("mysql", "root:mustang@tcp(localhost:3306)/test?timeout=90s&collation=utf8mb4_unicode_ci")
	if err != nil {
		fmt.Errorf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	insert(db)

	rows, err := db.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}