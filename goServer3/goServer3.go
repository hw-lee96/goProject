package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var userId string

	err = db.QueryRow("SELECT userId FROM user WHERE idx = 1").Scan(&userId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)
}
