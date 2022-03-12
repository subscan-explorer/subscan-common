package sql

import (
	"context"
	"fmt"
	"testing"
)

type Book struct {
	ID     int64
	Title  string
	Author string
	Date   string
}

// go test -v -test.run TestQueryRow
func TestQueryRow(t *testing.T) {
	sqlStr := "select * from test_tbl where id=?;"
	var book Book
	err := db.QureyRow(context.TODO(), sqlStr, 1).Scan(&book.ID, &book.Title, &book.Author, &book.Date)
	if err != nil {
		fmt.Println(err) // proper error handling instead of panic in your app
		return
	}
	fmt.Println(book)
}

// go test -v -test.run TestQuery
func TestQuery(t *testing.T) {
	sqlStr := "select * from test_tbl;"
	// Execute the query
	results, err := db.Qurey(context.TODO(), sqlStr)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var book Book
		// for each row, scan the result into our tag composite object
		err = results.Scan(&book.ID, &book.Title, &book.Author, &book.Date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(book)
	}
}
