package main

import "fmt"

func main() {

	//array
	ar := [4]int{1, 2, 3, 4}
	fmt.Println(ar)

	slices := []int{1, 2, 3}
	fmt.Println(slices)

	//appending to slices
	slices = append(slices, 8)

	fmt.Println(slices)

}
