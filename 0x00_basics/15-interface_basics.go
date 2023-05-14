package main

/**
*Interfaces in Go
*
* Interfaces group types together based on their methods
*
*English: interfaces are exclusive (members-only, for americans
*and americanised in the crowd) clubs
*to be "enrolled into" (be a member of) an interface, a struct, or
*equivalent must have, associated with its type, all
*the methods declared in the interface.
*
*For dummies: an interface is like a golf club; one must meet
*all criteria to be admitted.
 */

import (
	"fmt"
	"math"
)

//`shape` interface
type shape interface {
	area() float64
	//perimeter() float64
	circumference() float64
	//can these be optional? I dont know
}

//`square` and `circle` structs
type square struct {
	length float64
}
type circle struct {
	radius float64
}

//square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumference() float64 {
	return s.length * 4
}

//circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

//func to print out the info of the shape
//must be used by both `square` and `circle`
func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is %0.2f\n", s, s.area())
	fmt.Printf("Perimeter of %T is %0.2f\n", s, s.circumference())
	//fmt.Printf("Circumference of %T is %0.2f\n", s, s.circumference())
}

//driver
func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*-#*")
	}
}
