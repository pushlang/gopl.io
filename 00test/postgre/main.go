package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type courses struct {
	cNo   string
	title string
	hours int
}

func main() {

	connStr := "user=postgres password=12481 dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from courses")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []courses{}

	for rows.Next() {
		p := courses{}
		err := rows.Scan(&p.cNo, &p.title, &p.hours)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.cNo, p.title, p.hours)
	}
}
