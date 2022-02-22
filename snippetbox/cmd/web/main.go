package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/mohatb/snippetbox/pkg/handlers"
)

const portNumber = ":4000"

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	//get the ip of the machine

	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		errorLog.Println(error)

	}
	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)

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
		ErrorLog: errorLog,
		Handler:  mux,
		Addr:     portNumber,
	}

	infoLog.Printf("Starting server on http://%s%s", ipAddress.IP.String(), portNumber)
	// Call the ListenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
