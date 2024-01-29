package main

import "fmt"

func forSelectLoop()  {
	charChan := make(chan string, 5)
	charSlice := []string{"a", "b", "c", "d", "e"}

	for _, val:=range charSlice {
		// select{
		// case charChan <- val:
		// }
		charChan <- val
	}
	close(charChan)

	for val:=range charChan {
		fmt.Println(val)
	}
}