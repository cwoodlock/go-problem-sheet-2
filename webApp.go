//Web application in golang problem sheet 2
//Author: Colm Woodlock
//Student Number: G00341460
//Adapted from: https://gowebexamples.com/templates/ and https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
package main

import (
	"log"
	"net/http"
)

type Prompt struct {
	Prompt string
}

func main() {
	//Call the template handler
	http.HandleFunc("/guess", TemplateHandler)

	//Serve to 127.0.0.1:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Function for the remplate handler
func TemplateHandler(w http.ResponseWriter, r *http.Request) {

	//********TRIED BOTH OF THESE WAYS COULDNT GET THEM TO WORK (compiles for both)*********

	/*//Message to be used as the prompt
	m := Prompt{Prompt: "Guess a number between 1 and 20"}
	//parse from index.tmpl
	t, _ := template.ParseFiles("guess/index.tmpl")
	t.Execute(w, m)*/

	/*gTemplate := template.Must(template.ParseFiles("guess/index.tmpl"))
	//Merge
	gTemplate.Execute(w, Message{Message: "Guess a number between 1 and 20"})*/
}
