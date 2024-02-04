package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fanIn[T any](done <-chan int, primeFinderChans ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i:=range c {
			select {
			case <-done:
				return
			case fannedInStream <-i:
			}
		}
	}

	for _, c:=range primeFinderChans {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func runFanInFanOut()  {
	start := time.Now()

	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {return rand.Intn(1000000000)}
	randIntStream := generator(done, randNumFetcher)	
	
	//fan out
	CPUCount := runtime.NumCPU()
	primeFinderChans := make([]<-chan int, CPUCount)
	for i:=0; i<CPUCount; i++ {
		primeFinderChans[i] = primeFinder(done, randIntStream)
	}

	//fan in
	fannedInStream := fanIn(done, primeFinderChans...)
	for i:=range take(done, fannedInStream, 20) {
		fmt.Println(i)
	}

	fmt.Printf("fan-in-fan-out takes %v\n", time.Since(start))
}