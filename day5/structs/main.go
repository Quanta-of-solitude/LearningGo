package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson"}
	var robi person

	fmt.Println(alex)
	fmt.Println(robi)
	robi.firstName = "Robi"
	robi.lastName = "Romen"
	robi.contactInfo.email = "robi@ramen.me"

	fmt.Printf("%+v", robi) //%+v prints alll field names and values

	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contactInfo: contactInfo{
			email:   "jim@notme.jim",
			zipCode: 198212,
		},
	}
	fmt.Println(jim)
	//call the receiver function
	jim.print()
	//update jim's name
	jim.updateName("Jimson")
	jim.print()
	// we will now use pointers

	jimPointer := &jim //get the address of jim
	jimPointer.updateNameProper("Jimson")
	jimPointer.print()

	//shortut for the above
	//after creating a struct receiver function with pointer (our pointerToPerson *person)
	//we can do the same without storing the address as well like:
	//jim.updateNameProper("Jimmy") <- see the receiver function is actually pointer and Go Handles the reference for us
	//jim.print()
	//Will work the same
}

//receiver function for type person

func (p person) print() {
	fmt.Printf("%+v", p)

}

//update name receiver function

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// update with pointers to haev effect
func (pointerToPerson *person) updateNameProper(newName string) {
	(*pointerToPerson).firstName = newName
}
