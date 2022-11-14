// this is going to be inside our main package so package main
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings" //used to convert to string from the deck type in toString method of deck
	"time"
)

//creating a new type of string called deck (slice)

type deck []string

// this is returning of type deck which in actual is a slice of strings
// this does not belong to type deck, but we can call it in main by simply declaration
// and since it returns type deck we can later call print() on that declared var
// Example: (in main.go)
// mycards := newDeck()
// mycards.print()
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
			//this will create our card deck instead of typing them manually:
			//Ace of Spades, Two of Spades, Three of Spades ...
			//Ace of Diamonds, Two Diamonds .......
			//Ace of Hearts, Two of Hearts and so on...
		}
	}
	return cards //cards is of type deck
}

// a function that belongs to deck type and to be called with the instance of type deck after assignment
func (d deck) print() {

	for _, card := range d {
		fmt.Println(card)
	}
}

// function which accepts the card deck as d, and handsize (number of cards in hand) of int type
// In this funcyion we take the starting deck and split it up into two items, the number of cards in hands(which in hand)
// And the other is the card left in deck (not in hand)
// This function do not belong to deck type, it accepts deck. We can call this function directly in our main.go without declaration

func deal(d deck, handSize int) (deck, deck) {

	//d[:handSize] slicing the deck (slice) to return the cards in hand and d[handSize:] is not in hand
	return d[:handSize], d[handSize:]
}

//we will use ioutil package to save to file
// the deck msut be byte type for our Write File (io function)
//deck is of slice (string) typ, we have to convert the deck type to byte in the function
//it is like []byte(stringData)

// we will write this toString() just to convert to string
// this function belongs to type deck and must be called : deckInstance.toString()
// to string typecase: []string(otherDataType)
// here deck -> slice of string -> string datatype(Join)
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //take the slice of string convert to single string joined by commas
	//Ace of Spades,Two of Spades,Three of Spades,etc...
}

// saving to file using ioutil package
// this is a deck method and can be called using deckInstance.saveToFile(filename)
func (d deck) saveToFile(filename string) error { //return error message if there is an error from the io WriteFile
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //[]byte(d.toString()) <- convert to byte type from string, 0666 is permission

}

// read deck from a file
// we are not making this of type deck because when reading the deck doesnt exist by default
// return type is deck because after reading the deck, it must become a deck type
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename) //byteSlice is where the byte data from the read file is present
	// err is the return of error from the ReadFile() func (Documentation io)

	//if there is error:like file not existing, nil is None
	if err != nil {
		//log the error and create a newDeck instead using newDeck()
		//OR log the error and quit the program <- using this for now
		fmt.Println("Error: ", err)
		os.Exit(1) //exit the program
	}
	//if there was no error, we covert the byteSlice to string data
	//string(byteSlice) //converts from [0 76 121 ] kind of data to Ace of Spades,Two of Spades,Three of Spades,etc...
	//we need to convert it back to slice of string using Split() from string package
	s := strings.Split(string(byteSlice), ",") //converts the comma separated values to []string slice
	//convert the slice of string (s) back to deck type
	return deck(s)
}

// shuffle the deck
// this method belongs to deck so we can call deckInstance.shuffle()
func (d deck) shuffle() {
	//seed the random generator
	rand.Seed(time.Now().UnixNano())
	// we are only looking at indexes not the card values
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		//swap elements randomly based on index values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//we completed out basic card project
