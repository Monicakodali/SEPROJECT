package main

import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    fmt.Printf("Starting server at port 8080\n")
	// adding route handlers to web server
	// Argument 1: pathname
	// Argument 2: Function that holds business logic to respond to request
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello World!")
    })

	// Starting the web server
	// Argument 1: portno
	// Argument 2: handler to configure server for HTTP/2 (default:nil)
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
