package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("Server is running on port 8080")

	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/8ball", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "higgyhiggy time")
	})
	http.ListenAndServe(":8080", nil)
}
