package main

import "fmt"

func getLast[T any](s []T) T {
	if len(s)>0 {
		return s[len(s)-1]
	}

	var zeroVal T
	return zeroVal
}

type email struct {
	message string
	senderEmail string
	recipientEmail string
}

type payment struct {
	amount int
	senderEmail string
	recipientEmail string
}

func main()  {
	test([]email{}, "email")
	test([]email{
		{
			"Hi, goat matata",
			"dadangombe@example.com",
			"goatmatata@example.com",
		},
		{
			"Hi, dada ng'ombe",
			"dadangombe@example.com",
			"goatmatata@example.com",
		},
		{
			"can i promote Chizi About Cheese on your podcast?",
			"dadangombe@example.com",
			"goatmatata@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"mzeekobe@example.com",
			"bibuibui@example.com",
		},
		{
			25,
			"kakasungura@example.com",
			"pakathecat@example.com",
		},
		{
			1,
			"papatheshark@example.com",
			"kakambwamwitu@example.com",
		},
		{
			16,
			"kakambweha@example.com",
			"kakadubu@example.com",
		},
	}, "payment")
}

func test[T any](s []T, desc string)  {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length %v\n", desc, len(s))
	for i, v:=range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println("=====================================================")
}