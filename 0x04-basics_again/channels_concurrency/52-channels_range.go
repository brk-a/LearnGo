package main

import (
	"fmt"
	"time"
)

func concurrentFib(n int)  {
	ch := make(chan int)
	
	go func()  {
		fibonacci(n, ch)
	}()

	for i:=range ch {
		fmt.Println(i)
	}
}

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================================")
}

func main()  {
	test(10)
	test(-2)
	test(0)
	test(20)
	test(100)
}

func fibonacci(n int, ch chan int)  {
	x, y := 0, 1
	for i:=0; i<n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond*10)
	}
	close(ch)
}