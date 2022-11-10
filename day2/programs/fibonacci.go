package main

import "fmt"

func main() {

	a, b, c := 1, 1, 0

	for i := 0; i < 10; i++ {
		a = b
		b = c
		c = a + b
		fmt.Println(c)
	}

}
