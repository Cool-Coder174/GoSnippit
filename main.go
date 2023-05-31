package main

import (
	"log"
	"net/http"
)

// define a home handler for function which writes a byte slice containing "Hello from Snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches
	// the http.NotFound() functiom will send a 404 response to the client
	// then we return from the handler. If we don't return from the handler
	// if we do not return the handler then the code will keep exicuting and write the "Hello from snippet..." message

	if r.URL.Path != "/" {

		http.NotFound(w, r)
		return

	}

	w.Write([]byte("Hello from snippetbox"))

}

// add a snippit view handler
func snippitView(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from snippet"))

}

// add a sippet create handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a specific snippet..."))

}

func main() {

	// use the http.newServerMux() function to initialize a new server mux
	// then regester the home function as the handler for the '/' URL pattern

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippitView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// using the http.ListenandServe() function to start a new web server. We will pass in two paramaters.
	// The TCP network address to listen on (in this case ":4000") snf the servermux we just created.
	// if http.ListenAndServe() returns an error
	// we will use the log.Fatal() function to log the error message and exit.
	// NOTE: That ANY error returned by http.ListenAndServe() is always non-nil
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
