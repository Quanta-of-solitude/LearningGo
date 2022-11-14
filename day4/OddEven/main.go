package main

import "fmt"

func main() {
	//empty slice of int type
	numbers := []int{}
	//append numbers 0 to 10 into the slice
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println("Slice elements: ", numbers) // print the slice

	//iterate through the list elements (not index) and say odd or even
	for _, i := range numbers {
		if i%2 == 0 {
			fmt.Printf("\n%v is Even", i)
		} else {
			fmt.Printf("\n%v is Odd", i)
		}
	}
}
