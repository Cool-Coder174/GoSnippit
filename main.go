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

// add snippetView handler function
func snippetView(write http.ResponseWriter, read *http.Request) {

	write.Write([]byte("Display a specific SNipbox"))

}

// add a snippetCreate handler function
func snippetCreate(write http.ResponseWriter, read *http.Request) {

	// Use r.method to check wether request is using POST or not
	if read.Method != "POST" {

		// Use the Header.Set() function to add an "Allow: POST" header to the response header map.
		// The first parameter is the header name and the second parameter is the header value.
		write.Header().Set("Allow", " POST")

		// If POST is not used, use the w.WriteHeader() function to send a 405 status
		// code and use the w.Write() method to wrte a "method not allowed" response body.
		//We then return from the function so that the subsequent code is not excecuted.
		// write.WriteHeader(405)
		// write.Write([]byte("Method not Allowed"))
		// using the http.Error() function will run BOTH .WriteHeader() and .Write()
		// FORMAT: http.Error('var name for httpResponseWriter' , 'Desired text', status)
		// The functions are the same the only differance is that since we are passing our http.ResponseWriter to another function,
		// We are also sending a a response to the user for us.
		// also use http.'Status name" to avoid using literals.
		http.Error(write, "Method Not Allowed", http.StatusMethodNotAllowed)

		return

	}

	write.Write([]byte("Create new Snippet"))

}

func main() {

	// use the http.newServerMux() function to initialize a new server mux
	// then regester the home function as the handler for the '/' URL pattern

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
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
