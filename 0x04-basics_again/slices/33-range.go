package main

import "fmt"

func indexOfFirstBadWord(msg, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word==badWord {
				return i
			}
		}
	}

	return -1
}

func test(msg, badWords []string)  {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message [%v ]for bad words...", msg)

	for _, badWord := range badWords {
		fmt.Println(`- `, badWord)
		fmt.Printf("Index: %v\n", i)
		fmt.Println("================================================")
	}
}

func main()  {
	
	badWords := []string {"fuck", "shit"}
	msg := []string {"hi", "there", ",", "goat", "matata"}
	test(msg, badWords)

	badWords = []string {"fuck", "shit", "bloody"}
	msg = []string {"what", "the", "bloody", "fuck", "is", "flabbanaba", "?"}
	test(msg, badWords)

	badWords = []string {"fuck", "shit", "there"}
	msg = []string {"that", ",", "right", "there", ",", "is", "shit"}
	test(msg, badWords)

	badWords = []string {"fuck", "shit"}
	msg = []string {"", "quick", "", "fox", "", "over"}
	test(msg, badWords)
}