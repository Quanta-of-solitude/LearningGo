// go Routines and channels
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://nyrize.herokuapp.com",
		"http://google.com",
		"http://coditation.com",
		"http://facebook.com",
	}
	c := make(chan string) //make a channel

	for _, link := range links {
		//fmt.Println(link)
		//checkLink(link) //this is sync kind of code
		//introducing child routines
		go checkLinkwithChannel(link, c) //why does this not give output
		//use the go keyword and pass the channel
		//we need to use channels to communicate with the child routines
		//we also need to receive the channel value, else there will be no output
	}

	//fmt.Println(<-c) //print received value

	//using the above approach we get only the first link that is up. Because the child routine that runs first and gets returns first sends value
	//to the main routine and code exits! Other links are not looped or data is not reflect we solve this in the following way:
	//using a forloop to get the channel outputs according to the number of requests that we made

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	//using an alternative for loop syntax to check the links repeatedly

	//for {
	//	go checkLinkwithChannel(<-c,c)
	//}

	//we can do the same using :
	for l := range c {
		//go checkLinkwithChannel(l, c) //why does this give output
		//using function literal
		//alwys pass a value to child routine instead od using reference to solve the prblem belows
		go func(link string) { //passing link here makes a copy instead of same refernce to solve the below problem if using l directly
			time.Sleep(2 * time.Second)
			checkLinkwithChannel(link, c)
			//checkLinkwithChannel(l, c) //if we do not pass l as an arguement here, after the last link is received it will keep printing the last value
			//because the l is on outer scope, here we need to pass l as an arguemen tto the function literlat
		}(l) //like this
	}
}

func checkLinkwithChannel(link string, c chan string) { //we have to specify the channel return type as well
	//time.Sleep(2 * time.Second) //wait for 2 seconds
	_, err := http.Get(link) //this is a blocking call, which slows than our code, this is like sync code
	if err != nil {
		fmt.Printf("\n%v link is down\n", link)
		//c <- "Might be down!" //send this value to channel
		c <- link //to check the links repeatedly
		return
	}
	fmt.Printf("\n%v is up\n", link)
	//c <- "It is up!"
	c <- link //to check the links repeatedly
}

func checkLink(link string) {
	_, err := http.Get(link) //this is a blocking call, which slows than our code, this is like sync code
	if err != nil {
		fmt.Printf("\n%v link is down\n", link)

		return
	}
	fmt.Printf("\n%v is up\n", link)
}
