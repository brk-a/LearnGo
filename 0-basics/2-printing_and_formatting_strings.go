package main

//a few methods inside pkg fmt

import "fmt"

func main() {
	//Print does not add a newline; Println does
	fmt.Print("Using Print:")
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Println("Using Println:")
	fmt.Println("Hello, ")
	fmt.Println("World!")

	//Non-formatted strings
	fname := "Goat"
	lname := "Matata"
	fmt.Println("My name is", fname, lname, ".")

	//Formatted strings
	//use Printf(). The f stands for...you guessed it...
	//...formatted strings.
	// %v is default for all vars, regardless of type
	//%q places quote marks `" "` around strings and fucks up everything else
	//%T shows the type of a var
	//%f for floats. 6dp default. Specify dp to apply viz: %0.xf(x->no of dp)
	//Sprintf() saves the formatted str to a var once it has done the Printf magic
	//in other words, S[ave]Printf
	fname1 := "Dada"
	lname1 := "Ng'ombe"
	horns := 2
	horns_str := "2"
	place := "Afrika Mashariki"
	fmt.Printf("Hujambo! Jina langu ni %v %v.\n", fname1, lname1)
	fmt.Println("Using `percent q`...")
	fmt.Printf("Mimi ni mkaazi wa %q na nina pembe %q\n", place, horns)
	fmt.Println("change the second arg to type str...")
	fmt.Printf("Mimi ni mkaazi wa %q na nina pembe %q\n", place, horns_str)
	fmt.Printf("var horns is of type %T and var horns_str is of type %T\n", horns, horns_str)
	fmt.Printf("value of pi to 3dp is %0.3f\n", 3.14159)
	var str_f = fmt.Sprintf("%v %v and %v %v are very good friends, yaani marafiki wa %v", fname, lname, fname1, lname1, "dhati")
	fmt.Println(str_f)
}
