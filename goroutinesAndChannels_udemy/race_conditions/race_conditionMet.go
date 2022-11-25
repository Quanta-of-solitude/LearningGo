package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {

	msg = "Hello World"

	//in this code there is race condition. When we print out msg we are not sure which message will print out first
	wg.Add(2)
	go updateMessage("Hello universe!")
	go updateMessage("Hello cosmos!")

	wg.Wait()
	fmt.Println(msg) //try with go run -race main.go
	//we have no idea which message will finish first, hence wedont know the value of msg at the end of the file

}
