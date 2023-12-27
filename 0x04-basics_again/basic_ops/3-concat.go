package main

import (
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func main()  {
	fmt.Println(concat("Hi.", "My name is Goat Matata."))
	fmt.Println(concat("Jambo.", "Jina langu ni Dada Ng'ombe"))
}