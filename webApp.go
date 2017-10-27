//Web application in golanf problem sheet 2
//Author: Colm Woodlock
//Student Number: G00341460

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Main that starts a web application on port 8080
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}

//This prints Guessing game on tthe page 127.0.0.1:8080
func requestHandler(w http.ResponseWriter, r *http.Request) {
	//this ensures that the website reads in the text as HTML
	w.Header().Set("Content-Type", "text/html")
	//Added h1 tags
	fmt.Fprintf(w, "<h1>Guessing Game</h1>")

}
