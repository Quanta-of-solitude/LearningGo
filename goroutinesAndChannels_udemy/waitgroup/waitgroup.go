package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() //means one wait group is done
	fmt.Println(s)
}

func main() {
	//using wait groups to get proper output
	var wg sync.WaitGroup

	words := []string{"one", "two", "three"}
	wg.Add(len(words))

	for _, x := range words {
		go printSomething(x, &wg) //always pass waitgroups by reference
	}
	//the same can be obtained without waitgroups
	//we can use time.Sleep, but that is not a good option
	wg.Wait()

}
