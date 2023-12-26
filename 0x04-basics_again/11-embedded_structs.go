package main

import "fmt"

type sender struct {
	aUser
	rateLimit int
}

type aUser struct {
	name string
	number int
}

func testEmbedded(s sender)  {
	fmt.Println("sender's name: %v", s.name)
	fmt.Println("sender's number: %v", s.number)
	fmt.Println("sender's rateLimit: %v", s.rateLimit)
	fmt.Println("============================================================")
}

func main() {
	testEmbedded(sender{
		rateLimit: 10000,
		aUser: aUser{
			name:   "Kaka Mbweha",
			number: 254701234567,
		},
	})
	testEmbedded(sender{
		rateLimit: 5000,
		aUser: aUser{
			name:   "Dada Ng'ombe",
			number: 254712345678,
		},
	})
	testEmbedded(sender{
		rateLimit: 1000,
		aUser: aUser{
			name:   "Kaka Tai",
			number: 254723456789,
		},
	})
}