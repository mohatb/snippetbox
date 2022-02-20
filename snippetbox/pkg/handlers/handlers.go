package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mohatb/snippetbox/pkg/render"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts := render.RenderTemplates(w, "home.page.tmpl")
	err := ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	ts := render.RenderTemplates(w, "about.page.tmpl")
	err := ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// About is the about page handler
func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)
		//you must call w.WriteHeader() before any call to w.Write()

		//w.WriteHeader(405)
		//w.Write([]byte("Method Not allowed"))
		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.

		//below is a shortcut for the above.
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create Snippet"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {

	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.

	//It needs to retrieve the value of the id parameter from the URL query string, which we
	//can do using the r.URL.Query().Get() method. This will always return a string value for
	//a parameter, or the empty string "" if no matching parameter exists.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.

	//We used fprintf takes an io.writer. this is because it is an interface, whenever you see io.writer it is fine to pass it http.responsewriter. below is the fprintf source
	//func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	// this is why we passed it w which is a "http.responsewriter"

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
