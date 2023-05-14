package main

import "fmt"

//this file has no main() function because it will be imported as a module by 9-package_scope.go

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello,", n)
}

//all vars or funcs declared at the top level of this file are accessible,
// through `package main`, to any othe file in this dir
//As such, they are accesible from 9-package_scope.go
//Case applies vice versa

func showNum() {
	fmt.Printf("The number is: %v\n", num)
}
