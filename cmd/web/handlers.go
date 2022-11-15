package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a health check handler
func home(w http.ResponseWriter, r *http.Request) {
	// Checking if the URL patyh exactly matches "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Use template.ParseFiles() to read template file into a template set
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Use Execute() method on template set to write the template
	// content as response body
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	// Check if request is POST or not
	if r.Method != http.MethodPost {
		// Add 'Allow:POST' header to response
		w.Header().Set("Allow", http.MethodPost)

		// Send a 405 error if request is not POST
		http.Error(w, "Method Not Allowed", 405)
	}
	w.Write([]byte("Create a new snippet"))
}
