//Struct

package main

import "fmt"

type Name struct {
	firstName string
}

type Employee struct {
	name       Name
	age        int
	department string
}

func main() {
	cat := Employee{Name{"Josh"}, 21, "IT"}
	fmt.Println(cat.name)

	//Array of Employees
	employees := []Employee{{Name{"jon"}, 22, "CS"}, {Name{"Joe"}, 26, "HR"}, {Name{"Doe"}, 54, "SDE"}}
	fmt.Println(employees)

	//iterating one by one just the names
	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i].name.firstName)
	}

}
