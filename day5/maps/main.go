package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f4530",
	}

	fmt.Println(colors)

	//map declaration
	//var c map[string]string
	// c:= make(map[string]string)

	colors["blue"] = "34343f"
	fmt.Println(colors)

	c := make(map[string]int)
	c["someval"] = 21
	c["rr"] = 10
	fmt.Println(c)

	delete(c, "rr")
	fmt.Println(c)
	c["noevalx"] = 1023
	printMap(c)

}

func printMap(c map[string]int) {
	for key, value := range c {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
