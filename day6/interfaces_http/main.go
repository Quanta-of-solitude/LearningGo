package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(res)
	//not the best method, GoLang has better approaches
	//bs := make([]byte, 99999)
	//res.Body.Read(bs)
	//fmt.Println(string(bs))
	//in this we are taking the output of the res.Body and showing it on terminal: (Writer, sourceOfData)
	lw := logWriter{}
	io.Copy(lw, res.Body)
}

// Custom implementation of the Writer Interface instead of using GO's inbuild .-. No idea what is happening tbh
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
