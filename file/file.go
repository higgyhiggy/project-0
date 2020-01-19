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
	//here we see flag to set name is active to to choose a specifc name if not a random one is general one is given
	if name {
		fmt.Println(len(os.Args))
		//these condition statement are to handle the erros if they put the flag but dont put a name
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
	// n the name the file will have
	return n
}
