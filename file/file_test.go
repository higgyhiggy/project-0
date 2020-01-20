package file

import (
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	var name bool = true
	var n, testString string
	var args = [4]string{"testfile", "-file", "test", "/temp"}
	n = time.Now().Format(args[3]+"/jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + args[2]
	if name {
		fmt.Println(len(args), "made up array to count as args input from cmd")
		//these condition statement are to handle the erros if they put the flag but dont put a name
		if len(args) >= 4 {

			n = time.Now().Format(args[3]+"/jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + args[2]
			//fmt.Println(n, 3)
			//os.MkdirAll(os.Args[3], 0777)
			testString = n
			if testString != n {
				t.Error("testString and n should be the same there is a problem in naming the file")
			} else if n == "" {
				t.Errorf("string  is the value of empty meaning the file name was never created")

			}
		} else if len(args) == 3 {

			n = time.Now().Format("jan-02-2006") + " " + time.Now().Format("03:04:05 PM") + ": " + args[2]
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

}

func ExampleCreate() {
	n := Create(false)
	fmt.Println(n != "")
	//Output:true
}
