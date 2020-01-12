package main

import (
	"fmt"
	"os"
	"os/exec"
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
	s := cdate + ":" + file.Create(name, fileName)
	fmt.Println(s, user)

	f, _ := os.Create(cdate + ":" + fileName)
	n2, _ := f.WriteString("an eye for an eye makes the whole world blind.")

	fmt.Println(n2, "bytes written succesffully")

	out, _ := exec.Command("code", s).Output()
	fmt.Printf("the date is %s\n", out)
	//fmt.Println(fileName)
	//fmt.Println(file.Create(file, "day 1"))

}
