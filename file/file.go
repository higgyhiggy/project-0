// Package file is an file.go libary
package file

import (
	"os"
	"time"
)

//Create is a function that takes in a string and  creates a file or opens a file
// but also appending the  date it was create jan 8 , 2020 : "filename"
func Create(name bool) string {
	var n string

	if name {

		if len(os.Args) < 3 {
			n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + "quicknotes"
		} else {
			n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + os.Args[2]
		}

	} else {
		n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + "quicknotes"
	}
	return n
}
