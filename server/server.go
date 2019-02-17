package server

import (
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

// Page struct
type Page struct {
	Title string
	Body  []byte
}

// function to load a page (any page)
func loadPage(title string) (*Page, error) {
	// filename := title + ".html"
	// read file from view/filename.html
	// TODO: read from a real file like
	// body, err := ioutil.ReadFile("views/" + filename)

	body := []byte("<html><body><h1>My Index Page</h1><div>Some body text</div></body></html>")

	// if err != nil {
	// 	return nil, err
	// }

	// return a Page object with the page title and html body
	return &Page{Title: title, Body: body}, nil
}

// function to handle "view" requests
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// get the substring after "/view" to use as the page title
	title := r.URL.Path[len("/view/"):]

	// load page with the title requested
	p, err := loadPage(title)

	if err != nil {
		log.Println(err)
	}

	// parse the HTML body to string to write on the response
	// TODO: not really using the page title for anything lol
	fmt.Fprintf(w, string(p.Body))
}

// dummy Helllo World handler for the root path
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// StartServer starts the server and a minimal router
func StartServer() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
