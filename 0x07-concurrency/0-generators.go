package main

import (
	"fmt"
	"math/rand"
)

func generator[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <-fn():

			}
		}
	}()

	return stream
}

func runGenerator() {
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {return rand.Intn(1000000000)}
	for i:=range generator(done, randNumFetcher) {
		fmt.Println(i)
	}
}