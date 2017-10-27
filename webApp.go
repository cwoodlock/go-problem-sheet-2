//Web application in golang problem sheet 2
//Author: Colm Woodlock
//Student Number: G00341460
//Adapted from http://idealogylabs.com/blog/golang/2015/07/24/serving-static-files-with-golang-gorilla-mux.html

package main

import (
	"log"
	"net/http"
)

func main() {
	//This will serve the page using bootstrap to 127.0.0.1:8080/guess
	index := http.StripPrefix("/guess", http.FileServer(http.Dir("index/")))
	http.Handle("/guess", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
