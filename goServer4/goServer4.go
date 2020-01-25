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

	// 복수 Row를 갖는 SQL 쿼리
	var id string
	var name string
	// DB마다 다른 Placeholder 형식 사용
	// ex : MySql은 ? 를 사용하고, Oracle은 :val1, :val2 등을 사용하고, PostgreSQL은 $1, $2 등을 사용.
	rows, err := db.Query("SELECT userId, userNm FROM user where idx >= ?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name)
	}
}
