package mysql

import (
	"database/sql"
	"fmt"
)

// MustConnectDB returns a pointer to the MySQL database or panics.
func MustConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:P@ssw0rd@/snippetbox?parseTime=true")
	if err != nil {
		fmt.Println("ERROR:", err)
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
