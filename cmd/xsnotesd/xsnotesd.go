package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/higgyhiggy/project-0/config"
	_ "github.com/higgyhiggy/project-0/config"
)

var name string
var textt []byte

type Mynotes struct {
	Passing string
	Fname   string
}

var Ohtml = Mynotes{"", ""}

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
	http.HandleFunc("/8ball", hee)

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
	Ohtml.Passing = string(textt)
	Ohtml.Fname = string(config.Txtname)
	char, _, err = reader.ReadRune()
	if err != nil {
		check(err)
	}
	switch char {
	case '\n':

		break
	}

}

func hee(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/higgy.html"))
	tmpl.Execute(w, Ohtml)
}
