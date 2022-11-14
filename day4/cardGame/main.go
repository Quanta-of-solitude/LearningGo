package main

import "fmt"

func main() {
	cards := newDeck() //this is a function in deck.go, it doesnot belong to deck type but returns deck type
	fmt.Println("Original Deck: ", cards)

	//deal returns two values of type deck
	hand, remainingCards := deal(cards, 5) //calling the deal function from deck.go, with params: cards (which is a deck type and has our cards), and handsize

	//we can call hand.print() and remainingCards.print() because both are now of type deck
	//and print() is a function that belomgs to deck type
	fmt.Println("\nCards in Hand: ")
	hand.print()
	fmt.Println("\nRemaing Cards in Deck: ")
	remainingCards.print()
	cards.print() //since print belongs to deck type from deck.go

	//example to convert string type to byte type, this is type coversion
	//exampleStr := "Hello World"
	//fmt.Println([]byte(exampleStr))
	fmt.Printf("\nDeck converted to string type: %v\n\n", cards.toString()) //returns the comma printed strings
	fmt.Printf("\nSaving the deck\n")
	//Saving the deck to a local file
	cards.saveToFile("my_cards")

	//Loading the deck from file
	fmt.Printf("\nLoading saved Deck\n")
	cardsLoaded := newDeckFromFile("my_cards")
	//shuffle the cards
	cardsLoaded.shuffle()
	fmt.Printf("\nShuffling Cards\n\n")
	cardsLoaded.print()
}
