//Maps in Go

package main

import "fmt"

func main() {

	myMap := make(map[string]string)

	fmt.Println(myMap)

	myMap["name"] = "John Doe"
	myMap["address"] = "21 Street"

	fmt.Println(myMap)

	//Looping get key, value

	for key, value := range myMap {
		fmt.Println(key, value)
	}
	delete(myMap, "address")
	fmt.Println(myMap)
}
