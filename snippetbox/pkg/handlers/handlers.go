package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mohatb/snippetbox/pkg/mysql"
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

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		println(w, "createsnippet function error")
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := mysql.Insert(title, content, expires)
	if err != nil {
		println(w, "createsnippet function error")
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
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
