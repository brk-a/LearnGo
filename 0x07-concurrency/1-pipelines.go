package main

import (
	"fmt"
	"math/rand"
)

func take[T any, K any](done <-chan K, stream <-chan T, n  int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)

		for i:=0; i<n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func runTake() {
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {return rand.Intn(1000000000)}
	randNumStream := generator(done, randNumFetcher)
	for i:=range take(done, randNumStream, 20) {
		fmt.Println(i)
	}
}