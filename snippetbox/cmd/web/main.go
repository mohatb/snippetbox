package main

import (
	"database/sql"
	"flag"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohatb/snippetbox/pkg/logger"
)

const portNumber = ":4000"

func main() {

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	//get the ip of the machine
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		logger.ErrorLog.Println(error)

	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)

	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "root:P@ssw0rd@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	//dsn := "web:pass@/" + ipAddress.IP.String() + "/snippetbox?parseTime=true"

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()
	//starting the web server
	//below is the old code before adding the new server
	// infoLog.Println("Starting server on http://" + ipAddress.IP.String() + portNumber)
	// err := http.ListenAndServe(portNumber, mux)
	// errorLog.Fatal(err)

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		ErrorLog: logger.ErrorLog,
		Handler:  routes(), // Call the new routes.go routes() method and get mux
		Addr:     portNumber,
	}

	logger.InfoLog.Printf("Starting server on http://%s%s", ipAddress.IP.String(), portNumber)
	// Call the ListenAndServe() method on our new http.Server struct.
	err = srv.ListenAndServe()
	logger.ErrorLog.Fatal(err)

}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
