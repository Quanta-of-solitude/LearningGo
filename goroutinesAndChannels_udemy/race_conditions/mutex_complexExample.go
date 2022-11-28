package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//var for bank balance

	var bankBalance int
	var balance sync.Mutex

	fmt.Println("Inital bank balance: ", bankBalance)

	//define weekly revenue
	income := []Income{
		{Source: "Main Job", Amount: 400},
		{Source: "Donation", Amount: 100},
		{Source: "Part-time", Amount: 200},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(income))

	for i, income := range income {
		//fmt.Println(i, income)
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d you earned $%d from source %s\n", week, income.Amount, income.Source)

			}
		}(i, income)
	}

	wg.Wait()

	fmt.Print("Your final bank balance is $", bankBalance)
}
