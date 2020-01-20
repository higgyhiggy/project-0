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
	fmt.Println("this client can only work if the server is running otherwise it wont.")
	fmt.Println(string(body))

}
