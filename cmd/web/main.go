package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define struct to hold application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Command line flag to set address at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Spin up a servermux and registry home func as a handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create file server to serve out files from static directory
	filerServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", filerServer))

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Spin up a web server by passing in TCP network and servermux
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
