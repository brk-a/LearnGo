package main

import (
	"fmt"
	"math/rand"
	"time"
)

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	primes := make(chan int)
	isPrime := func (randInt int) bool {
		for i:=randInt-1; i>1; i-- {
			if randInt%i==0 {
				return false
			}
		}

		return true
	}

	go func() {
		defer close(primes)

		select {
		case <-done:
			return
		case randInt:= <-randIntStream:
			if isPrime(randInt) {
				primes <-randInt
			}
		}
	}()

	return primes
}

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
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {return rand.Intn(1000000000)}
	randNumStream := generator(done, randNumFetcher)
	primesStream := primeFinder(done, randNumStream)
	for i:=range take(done, primesStream, 20) {
		fmt.Println(i)
	}

	fmt.Printf("naive approach takes %v\n", time.Since(start))
}