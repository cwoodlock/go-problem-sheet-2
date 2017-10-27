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
	fmt.Fprintln(w, "Guessing game")
}
