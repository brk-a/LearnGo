package main

import "fmt"

func pipeline()  {
	nums := []int{1, 2, 3, 4, 5}

	dataChan := sliceToChannel(nums)

	finalChan := sq(dataChan)

	for n:= range finalChan {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n:=range nums {
			// select {
			// case out <-n:
			// }
			out <-n
		}
		close(out)
	}()

	return out
}

func sq(dataChan <-chan int) <-chan int {
	sqChan := make(chan int)
	go func() {
		for n:=range dataChan {
			// s := n*n
			// select {
			// case sqChan <-s:
			// }
			sqChan <-n*n
		}
		close(sqChan)
	}()

	return sqChan
}