//Pallindrome

package main

import "fmt"

func reverse(mystr string) string {
	r := ""
	for _, chr := range mystr {
		r = string(chr) + r
	}
	return r
}

func main() {
	strn := "racecar"
	value := reverse(strn)
	if strn == value {
		fmt.Println("Pallindrome")
	} else {
		fmt.Println("Not Pallindrome")
	}
}
