// Channels
// channels are more powerful than mutex and waitgroups
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error //channel of channels -> Here, a channel of channels that is holding errors
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber += 1
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1 //random number
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed += 1
		} else {
			pizzasMade += 1
		}
		total += 1
		fmt.Printf("Making Pizza$%d, will take %d seconds!\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("***Ran out of ingrediants for pizza #%d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("***Cook quit for pizza #%d\n", pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("***Pizza #%d is ready!\n", pizzaNumber)

		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzaeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0
	//run forever or until quit notifs
	//try to make pizzas

	for {
		//try to make pizzas
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			//we tried to make a pizza (we sent soemthing to the data chnnel)
			case pizzaMaker.data <- *currentPizza: //sending a chan PizzaOrder to the dataChannel of producer
			case quitChan := <-pizzaMaker.quit:
				//close our channels
				close(pizzaMaker.data)
				close(quitChan)

				//get out of goroutine
				return
			}
		}

	}

}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print out message
	color.Cyan("The Pizzaeria is open for business!")
	color.Cyan("-----------------------------------")

	//create producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	//run producer as background
	go pizzaeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad")
			}
		} else {
			color.Cyan("Done making Pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("** Error closing the channel", err)
			}
		}
	}
	//printout ending message
}
