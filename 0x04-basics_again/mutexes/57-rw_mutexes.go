package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux *sync.RWMutex
}

func (sc safeCounter) inc(key string)  {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string)  {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

func test(sc safeCounter, emailTests []emailTest)  {
	emails := make(map[string]struct{})
	var wg sync.WaitGroup

	for _, emailT:=range emailTests {
		emails[emailT.email] = struct{}{}
		for i:=0; i<emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest){
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email:=range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	sc.mux.RLock()
	defer sc.mux.RUnlock()
	for _, email:=range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=============================================================")
}

func main()  {
	sc := safeCounter{
		counts: make(map[string]int),
		mux: &sync.RWMutex{},
	}
	test(sc, []emailTest{
		{
			email: "goatmatata@example.co.ke",
			count: 23,
		},
		{
			email: "dadangombe@example.co.ke",
			count: 29,
		},
		{
			email: "kakasungura@example.co.ke",
			count: 31,
		},
		{
			email: "mzeekobe@example.co.ke",
			count: 67,
		},
	})
	test(sc, []emailTest{
		{
			email: "kakatai@example.co.ke",
			count: 23,
		},
		{
			email: "bibuibui@example.co.ke",
			count: 126,
		},
		{
			email: "kakambwamwitu@example.co.ke",
			count: 31,
		},
		{
			email: "kakambweha@example.co.ke",
			count: 453,
		},
	})
}