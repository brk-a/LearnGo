package main

//Code that executes only when certain condition(s) are met

import "fmt"

func main() {
	//A few expressions whose result is a boolean value (T|F)
	num := 45

	fmt.Println(num <= 50)
	fmt.Println(num >= 50)
	fmt.Println(num == 45)
	fmt.Println(num != 50)

	//The if statement executes code IFF the condition evaluates to T
	if num < 30 {
		fmt.Println("The number is less than 30")
	} else if num < 40 {
		fmt.Println("The number is less than 40, but greater than 30")
	} else if num < 50 {
		fmt.Println("The number is less than 50 but greater than 40")
	} else {
		fmt.Println("The number is greater than 50")
	}

	//the `continue ` and `break` statements...
	animals := []string{"Goat_Matata", "Dada_Ng'ombe", "Kaka_Mbweha", "Mzee_Kobe", "Kaka_Mbwamwitu", "Kaka_Dubu", "Paka_the_Cat", "Kaka_Sungura"}
	fmt.Println("Original slice:", animals, "Length:", len(animals))
	for index, value := range animals {
		if index == 3 {
			fmt.Println("Continue statement at index:", index, "whose value is:", value)
			continue
		}
		if index > 5 {
			fmt.Println("Break statement at index:", index, "whose value is:", value)
			break
		}
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

}
