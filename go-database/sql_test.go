package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	
	script := "INSERT INTO customer(id, name) VALUES(3, 'John')"

	_, err := db.ExecContext(ctx, script)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data")
}

func TestRowSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	
	script := "SELECT id, name FROM customer;"

	rows, err := db.QueryContext(ctx, script)
	
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
		fmt.Println()
	}

	defer rows.Close()
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	ctx := context.Background()
	
	script := "SELECT id, name, email, balance, birthdate, rating, married, created_at FROM customer"
	
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, balance int32
		var name string
		var email sql.NullString
		var created_at time.Time
		var birthdate sql.NullTime
		var rating float64
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &birthdate, &rating, &married, &created_at )
		if err != nil {
			panic(err)
		}

		fmt.Println("=============================================")
		fmt.Println("id         : ", id)
		fmt.Println("name       : ", name)
		if email.Valid { 
			fmt.Println("email      : ", email.String) 
		} 
		if birthdate.Valid { 
			fmt.Println("birthdate  : ", birthdate.Time.String()) 
		} 
		fmt.Println("balance    : ", balance)
		fmt.Println("rating     : ", rating)
		fmt.Println("married    : ", married)
		fmt.Println("created_at : ", created_at)			

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	username := "baihaqi16 '; #"
	password := "12345678"

	script := "SELECT username FROM users WHERE username = '"+ username +"' AND password = '"+ password +"';"

	rows, err := db.QueryContext(ctx, script)
	
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login", username)
	}else{
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	username := "baihaqi16"
	password := "12345678"

	script := "SELECT username FROM users WHERE username = ? AND password = ? ;"
	rows, err := db.QueryContext(ctx, script, username, password)
	
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success Login", username)
	}else{
		fmt.Println("Login Failed")
	}
}

func TestInsertSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	id := 4
	username := "admin'; DROP TABLE users; #"
	password := "admin"
	
	script := "INSERT INTO users(id, username, password) VALUES(?, ?, ?)"

	_, err := db.ExecContext(ctx, script, id, username, password)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	
	email := "yazid@gmail.com"
	comment := "Test Comment2"
	
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Comment With ID", id)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "yazid" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment "+ strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID", id)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 1; i <= 10; i++ {
		email := "baihaqi" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment "+ strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}