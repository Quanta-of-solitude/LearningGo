package main

import (
	"fmt"
	"log"
	"net/http"
)

// function that handles the hello route
// w -> Response of that route (sent to user) r-> is what sent by the user
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { //check if URL path is correct, we already handled it in main, but trying it out
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not allowed", http.StatusNotFound)
		return
	}
	//when on /hello send "hello!" to user's browser
	fmt.Fprintf(w, "hello!")
}

// function that handles the form route action
// w -> Response of that route (sent to user) r-> is what sent by the user
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { //r.ParseForm() is to get data what user types in the form (see it's a pointer type)
		fmt.Fprintf(w, "ParseForm() err %v", err)
	}
	fmt.Fprintf(w, "POST request successful")
	fname := r.FormValue("fname") //get the fname field value from form.html
	lname := r.FormValue("lname") //get the lname field value from form.html

	//after submssion we are Printing this to the user's browser
	//we can do a lot of different things here
	fmt.Fprintf(w, "FIRST Name = %s\n", fname)
	fmt.Fprintf(w, "LAST Name = %s\n", lname)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) //define the static directory, and by defualt it directs to the index.html
	http.Handle("/", fileServer)                        //direct the / route to index.html
	http.HandleFunc("/form", formHandler)               //for /form route invoke the formHandler Function
	http.HandleFunc("/hello", helloHandler)             //for the /hello route invoke the helloHandler Function

	fmt.Println("Starting server at port 8080")

	//http.ListenAndServer serves the webserver at PORT 8080
	if err := http.ListenAndServe(":8080", nil); err != nil { //we are checking error as well
		log.Fatal(err)
	}
}
