//Guessing Game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func get_guess() int {
	var res int

	fmt.Scan(&res)
	return res
}

func main() {

	rand.Seed(time.Now().Unix())
	var generated int
	//fmt.Println(generated)
	//fmt.Println("Give your guess: 0-10 or 11 to exit")
	val := 12

	for val != generated {

		generated := rand.Intn(11)
		fmt.Println("\nGive your guess: 0-10 or 11 to exit")
		val = get_guess()
		if val == 11 {
			break
		}
		if val == generated {
			fmt.Println("\nCorrect!")
		} else {
			fmt.Println("\nIncorrect, number was: ", generated)
		}

	}

}
