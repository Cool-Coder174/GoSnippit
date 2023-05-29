package main

import (
	"log"
	"net/http"
)

// define a home handler for function which writes a byte slice containing "Hello from Snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from snippetbox"))

}

// add snippetView handler function
func snippetView(write http.ResponseWriter, read *http.Request) {

	write.Write([]byte("Display a specific SNipbox"))

}

// add a snippetCreate handler function
func snippetCreate(write http.ResponseWriter, read *http.Request) {

	write.Write([]byte("Create new Snippet"))

}

func main() {

	// use the http.newServerMux() function to initialize a new server mux
	// then regester the home function as the handler for the '/' URL pattern

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// using the http.ListenandServe() function to start a new web server. We will pass in two paramaters.
	// The TCP network address to listen on (in this case ":4000") snf the servermux we just created.
	// if http.ListenAndServe() returns an error
	// we will use the log.Fatal() function to log the error message and exit.
	// NOTE: That ANY error returned by http.ListenAndServe() is always non-nil
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
