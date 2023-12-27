package main

import (
	"errors"
	"fmt"
)


func main(){
	testDiv(2, 0)
	testDiv(3, 1)
	testDiv(10, 0)
	testDiv(7, 5)
	testDiv(1, 2)
}

func testDiv(x int, y int) {
	quotient, err := divide(x, y)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println("%v divide by %v is %v", x, y, quotient)
	}
	
	fmt.Println("========================================================")
}

func divide(dividend, divisor int) (int, error) {
	if divisor==0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return dividend/divisor, nil
}