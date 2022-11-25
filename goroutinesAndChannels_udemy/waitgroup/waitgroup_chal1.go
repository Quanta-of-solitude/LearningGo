package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	var wg sync.WaitGroup

	msg = "Hello world"

	//msgs := []string{"Hello, universe", "Hello cosmos", "Hello World"}
	/*wg.Add(len(msgs))

	for _, x := range msgs {
		go updateMessage(x, &wg)
		printMessage()
	}
	wg.Wait()*/
	wg.Add(1)
	go updateMessage("Hello Universe", &wg)
	printMessage()
	wg.Wait()

	wg.Add(1)
	go updateMessage("Hello Cosmos", &wg)
	printMessage()
	wg.Wait()

	wg.Add(1)
	go updateMessage("Hello World", &wg)
	printMessage()
	wg.Wait()

}
