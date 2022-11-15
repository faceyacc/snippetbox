package main

import (
	"log"
	"net/http"
)

func main() {

	// Spin up a servermux and registry home func as a handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create file server to serve out files from static directory
	filerServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	// Spin up a web server by passing in TCP network and servermux
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
