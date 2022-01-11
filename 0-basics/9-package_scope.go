package main

//scope of variables

import "fmt"

var num = 99.5

func main() {
	sayHello("Kaka Mbweha")

	for _, val := range points {
		fmt.Println(val)
	}

	showNum()
}

//Go is quite the language...
//Ignore those errors.
//In order to run the code successfully, do this...
// go run 9-package_scope.go 9-package_scope_helper.go
