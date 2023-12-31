package main

import (
	"errors"
	"fmt"
)

func getMessageCost(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i:=0; i<len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * .01
		costs[i] = cost
	}

	return costs
}

func test(messages []string) {
	costs := getMessageCost(messages)

	fmt.Println("Messages...")
	for i:=0; i<len(messages); i++ {
		fmt.Printf(`- %v\n`, messages[i])
	}

	fmt.Println("Costs...")
	for i:=0; i<len(messages); i++ {
		fmt.Printf(`- %v\n`, costs[i])
	}
}

func main() {
	test([]string{
		"",
		"The quick, brown fox jumps over the lazy dog",
	})
	test([]string{
		"dinglebop",
		"schleem",
		"grumbo",
		"fleeb",
		"schlami",
		"hizzards",
		"blampfs",
		"chumbles",
		"ploobis",
		"plumbus",
	})
	test([]string{
		"Everyone has a plumbus in their home. First, they take the dinglebop and they smooth it out with a bunch of schleem. The schleem is then repurposed for later batches. They take the dinglebop and they push it through the grumbo, where the fleeb is rubbed against it. It’s important that the fleeb is rubbed because the fleeb has all of the fleeb juice. Then a schlami shows up and he rubs it and spits on it. They cut the fleeb. There’s several hizzards in the way. The blamfs rub against the chumbles. And the ploobis and grumbo are shaved away. That leaves you with a regular old plumbus. ",
	})
}