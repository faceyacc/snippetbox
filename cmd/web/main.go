package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Command line flag to set address at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Spin up a servermux and registry home func as a handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create file server to serve out files from static directory
	filerServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	// Spin up a web server by passing in TCP network and servermux
	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
