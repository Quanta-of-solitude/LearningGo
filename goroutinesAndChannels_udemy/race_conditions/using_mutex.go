package main

// we are getting rid of race condition using mutex
//we are performing thread safe, and accessing data safely
//testing for race conditions is mandatory

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {

	msg = "Hello World"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello universe!", &mutex)
	go updateMessage("Hello cosmos!", &mutex)

	wg.Wait()
	fmt.Println(msg)
}
