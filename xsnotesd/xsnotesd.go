package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/higgyhiggy/project-0/config"
	_ "github.com/higgyhiggy/project-0/config"
)

var name string

func main() {

	fmt.Println("Server is running on port 8080")
	//takes a patter string for the handle and takes in a handler
	//then handler returns the content of a file system
	http.Handle("/", http.FileServer(http.Dir("web")))
	//deals with the handl with the anonymous handler function
	http.HandleFunc("/8ball", func(w http.ResponseWriter, r *http.Request) {
		//r is a request to url
		//we are working with form to be able to use user input
		//if no name is given then fall back valur will be anonymous
		name = r.FormValue("name")
		if name == "" {
			name = "anonymous"
		}
		fmt.Fprint(w, "<h1>"+config.Word[rand.Intn(len(config.Word))].Text+"</h1>", " ", name)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

	})
	// can make some infinite loop to add some other functionality
	go http.ListenAndServe(":8080", nil)

	time.Sleep(30 * time.Second)

}
