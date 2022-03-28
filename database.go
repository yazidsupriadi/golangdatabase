package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajargolangdatabase?parseTime=true")
	if err != nil {
		panic(err)
	}

	//pooling
	db.SetMaxIdleConns(10)                  //maksimal idle atau maksimal connection yang diam
	db.SetMaxOpenConns(100)                 //maksimal connection yang terbuka
	db.SetConnMaxIdleTime(5 * time.Minute)  //maksimal waktu idle
	db.SetConnMaxLifetime(60 * time.Minute) //maksimal lifetime db

	return db
}
