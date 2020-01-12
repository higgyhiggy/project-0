package main

import (
	"fmt"
	"os"

	"github.com/higgyhiggy/project-0/file"
)

func main() {
	fmt.Println("this is going to be a note taking app")
	fmt.Println("sho rururururururur")
	//asks the env to get user name
	user := os.Getenv("USER")
	fmt.Println(file.Create(name, "hey"), user)

	//fmt.Println(file.Create(file, "day 1"))

}
