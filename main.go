package main

import (
	"fmt"
	"os"
	"time"

	"github.com/higgyhiggy/project-0/file"
)

func main() {
	currentTime := time.Now()
	cdate := currentTime.Format("jan-02-2006")
	//asks the env to get user name

	//0 caller location path
	// 1 first arg
	//2 second arg and so on
	user := os.Getenv("USER")
	fileName := os.Args[2]
	fmt.Println(cdate+":"+file.Create(name, fileName), user)
	//fmt.Println(fileName)
	//fmt.Println(file.Create(file, "day 1"))

}
