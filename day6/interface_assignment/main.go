/*
Using interfaces to Print Area of Square and Triangle with a singular PrintFunction instead of using two different declarations
*/

package main

import "fmt"

// shape interface with getArea() that returns a float64 type
type shape interface {
	getArea() float64
}

// Struct of Traingle with field height and base of float64 type
type triangle struct {
	height, base float64
}

// Struct of Square with field sideLength of float64 type
type square struct {
	sideLength float64
}

func main() {
	sq := square{10.1}                       //calling the square struct with side value
	tr := triangle{height: 10.0, base: 12.0} //callimg the triangle struct with height and base values

	//Printing the area of square and triangle through shape interface function
	printArea(sq)
	printArea(tr)
}

// Interface function that takes in shape parameter and access the getArea() function
func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

// Calucate the area of Square which returns a value of type float64
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// Calculate teh area of Triangle which returns a value of type float64
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
