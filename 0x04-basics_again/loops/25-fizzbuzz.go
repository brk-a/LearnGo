package main

import "fmt"

func fizzbuzz(a int)  {
	i := 1
	for i<=a {
		if i%3==0 && i%5==0 {
			fmt.Println("fizzbuzz")
		} else if i%3==0 {
			fmt.Println("fizz")
		} else if i%5==0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

		i++
	}
}

func test(a int) {
	fmt.Printf("fizzbuzz for input %v\n", a)
	fizzbuzz(a)
	fmt.Println("=================================================================")
}

func main(){
	test(100)
	test(5)
	test(243)
}