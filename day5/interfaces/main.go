package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
////	fmt.Println(eb.getGreeting())
//}

// we can reuse this else we have to declar with different name
// both the functions are using same stuff so we can use interfaces to solve these problems
//func printGreeting(sb englishBot) {
//	fmt.Println(sb.getGreeting())
//}

func (eb englishBot) getGreeting() string {
	return "Hi!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
