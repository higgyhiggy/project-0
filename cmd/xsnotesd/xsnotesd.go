package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/higgyhiggy/project-0/config"
	_ "github.com/higgyhiggy/project-0/config"
)

var name string
var textt []byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}
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

		fmt.Fprintf(w, string(textt))

	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

	})

	go http.ListenAndServe(":8080", nil)
	// to be able to out put file context to the terminal
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		check(err)
	}

	switch char {
	case '\n':
		cmd := exec.Command("cat", config.Txtname)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			check(err)
		}
		break
	}

	//to out put file context to http
	content, err := ioutil.ReadFile(config.Txtname)
	if err != nil {
		check(err)
	}
	textt = content

	char, _, err = reader.ReadRune()
	if err != nil {
		check(err)
	}
	switch char {
	case '\n':

		break
	}

}
