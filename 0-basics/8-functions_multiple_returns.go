package main

import (
	"fmt"
	"strings"
)

//functions with multiple return values

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, val := range names {
		initials = append(initials, val[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""
}

func main() {
	fn1, ln1 := getInitials("Goat Matata")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("Dada Ng'ombe")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("Paka_the_Cat")
	fmt.Println(fn3, ln3)

	fn4, ln4 := getInitials("Paka the Cat")
	fmt.Println(fn4, ln4)
}
