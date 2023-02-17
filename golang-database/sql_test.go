package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `INSERT INTO customer(id, name) VALUES("adi", "Adi")`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name FROM customer`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("===================================")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date :", birth_date.Time)
		}
		fmt.Println("Created At :", created_at)
		fmt.Println("Married :", married)
		fmt.Println("===================================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "bike"

	query := "SELECT username FROM user WHERE username ='" + username + "' AND password='" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := `SELECT username FROM user WHERE username =? AND password=? LIMIT 1`
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecSqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	id := "awikwok"
	name := "awikwok"
	query := `INSERT INTO user(username, password) VALUES(?, ?)`
	_, err := db.ExecContext(ctx, query, id, name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "adi@anto.com"
	comment := "Halo nama saya bilek"

	query := `INSERT INTO comments(email, comment) VALUES(?, ?)`
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `INSERT INTO comments(email, comment) VALUES(?,?)`
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 1; i <= 10; i++ {
		email := "adi" + strconv.Itoa(i) + "siswanto@gmail.com"
		comment := "ini komen ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id ke-", id)
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

	query := `INSERT INTO comments(email, comment) VALUES(?, ?)`
	// do transaction
	for i := 1; i <= 10; i++ {
		email := "adi" + strconv.Itoa(i) + "siswanto@gmail.com"
		comment := "ini komen ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id ke-", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
