package main

import (
	"database/sql"
	"fmt"

	dynamicq "github.com/iivkis/dynamic-query"
)

func main() {
	// your database
	db := new(sql.DB)

	// with not nil param
	{
		age := dynamicq.PtrInt(18)

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		if age != nil {
			dq.Where("age = ?", age)
		}

		//paste "WHERE" to the query
		dq.Glue(&query)

		fmt.Println(query)
		/*
			output: SELECT name FROM user WHERE age = ?
		*/

		// executing a query with substitution of arguments
		db.Exec(query, dq.Args()...)
	}

	// with nil param
	{
		var age *int //nil pointer

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		if age != nil {
			dq.Where("age = ?", age)
		}

		dq.Glue(&query)
		fmt.Println(query)
		/*
			output: SELECT name FROM user
		*/

		db.Exec(query, dq.Args()...)
	}

	// adding attributes
	{
		age := 18 //nil pointer

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		dq.Where("age = ?", age)

		// paste "WHERE" to the query
		dq.Glue(&query)

		// add attribute
		dq.Attr(&query, "ORDER BY id")

		fmt.Println(query)
		/*
			output: SELECT name FROM user WHERE age = ? ORDER BY id
		*/

		db.Exec(query, dq.Args()...)
	}

	// limit
	{
		age := 18 //nil pointer

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		dq.Where("age = ?", age)

		dq.Glue(&query)
		dq.Attr(&query, "ORDER BY id")

		//does not add when the parameter is zero
		dq.Limit(&query, 0)
		fmt.Println(query)
		/*
		 output: SELECT name FROM user WHERE age = ? ORDER BY id
		*/

		dq.Limit(&query, 15)
		fmt.Println(query)
		/*
		 output: SELECT name FROM user WHERE age = ? ORDER BY id LIMIT 15
		*/

		db.Exec(query, dq.Args()...)
	}

	// offset
	{
		age := 18 //nil pointer

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		dq.Where("age = ?", age)

		dq.Glue(&query)
		dq.Attr(&query, "ORDER BY id")

		//does not add when the parameter is zero
		dq.Offset(&query, 0)
		fmt.Println(query)
		/*
		 output: SELECT name FROM user WHERE age = ? ORDER BY id
		*/

		dq.Offset(&query, 100)
		fmt.Println(query)
		/*
			output: SELECT name FROM user WHERE age = ? ORDER BY id OFFSET 100
		*/

		db.Exec(query, dq.Args()...)
	}

	// OR
	{
		age := 18 //nil pointer
		gender := "male"

		var (
			dq    = dynamicq.Dynamic{}
			query = "SELECT name FROM user"
		)

		dq.Where("age = ? OR gender = ?", age, gender)
		dq.Glue(&query)

		fmt.Println(query)
		/*
			output: SELECT name FROM user WHERE age = ? OR gender = ?
		*/

		db.Exec(query, dq.Args()...)
	}
}
