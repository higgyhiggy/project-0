package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//makes a requst to the 8ball handle
	//read the whole request body into our body variable
	// displays it in the terminal
	resp, _ := http.Get("http://localhost:8080/8ball")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println("this client can only work if the server is running otherwise it wont!!!!!\n\n\n")
	//response := body[892:1059]
	var i int
	var v byte
	prequest := body[916:]
	for i, v = range prequest {

		if v == 60 {
			break
		}

	}

	//fmt.Println(string(body[892:1059]))
	fmt.Print(string(prequest[:i]))
}
