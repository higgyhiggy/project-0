// Package file is an file.go libary
package file

import (
	"fmt"
	"os"
	"time"
)

//Create if a fucntion can create a file at a specified location
//it has some default fall backs if the file name is missing or if the path is missing as well
// and if the flag is missing also
func Create(name bool) string {
	var n string

	if name {
		fmt.Println(len(os.Args))
		if len(os.Args) >= 4 {

			n = time.Now().Format(os.Args[3]+"/jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + os.Args[2]
			//fmt.Println(n, 3)
			os.MkdirAll(os.Args[3], 0777)
		} else if len(os.Args) == 3 {

			n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + os.Args[2]
			//fmt.Println(n, 2)
		} else {
			n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + "quicknotes"
			//fmt.Println(n, 0)
		}

	} else {
		n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + "quicknotes"
		//fmt.Println(n, 0)
	}
	//return n

	//fmt.Println(n)
	return n
}
