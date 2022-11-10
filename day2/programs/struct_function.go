//USing a function that belongs to struct

package main

import "fmt"

type Triangle struct {
	base   float32
	height float32
}

func (areaTriangle Triangle) area() float32 {
	return 0.5 * areaTriangle.base * areaTriangle.height
}

func main() {

	triangle := Triangle{10, 5}
	fmt.Printf("Area of Triangle : %.2f \n", triangle.area())
}
