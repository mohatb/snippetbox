package mysql

import (
	"database/sql"
	"fmt"

	"github.com/mohatb/snippetbox/pkg/models"
	// Import the models package that we just created. You need to prefix this with
	// whatever module path you set up back in chapter 02.02 (Project Setup and Enabling
	// Modules) so that the import statement looks like this:
	// "{your-module-path}/pkg/models".
)

var DB = MustConnectDB()

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

func Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content and expiry values for the placeholder parameters. This
	// method returns a sql.Result object, which contains some basic
	// information about what happened when the statement was executed.
	result, err := DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// This will return a specific snippet based on its id.
func Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func Latest() ([]*models.Snippet, error) {
	return nil, nil
}
