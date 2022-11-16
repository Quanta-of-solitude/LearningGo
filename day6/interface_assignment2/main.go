// Read file from local and print it to terminal

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fmt.Println(os.Args[1])
	//got from documentation os.Open and os.Args
	myFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	//fmt.Println(myFile)
	//The File type implements the use of Reader Interface so io.Copy is usable
	//os.Stdout is the destination where to print contents of our file
	io.Copy(os.Stdout, myFile)
}
