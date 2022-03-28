package golangdatabase

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer VALUES ('2','yazid Supriadi')"

	_, err := db.ExecContext(ctx, script) //Exec.Content for insert,update,delete
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Customer to database Success")
}

func TestExecQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"

	rows, err := db.QueryContext(ctx, script) //Exec.Content for insert,update,delete
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
		fmt.Println("id", id)
		fmt.Println("name", name)
	}

}

func TestExecQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"

	rows, err := db.QueryContext(ctx, script) //Exec.Content for insert,update,delete
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float32
		var birth_date, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)

		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
		fmt.Println("balace", balance)
		fmt.Println("rating", rating)
		fmt.Println("birthdate", birth_date)
		fmt.Println("married", married)
		fmt.Println("created at", created_at)

	}

}
