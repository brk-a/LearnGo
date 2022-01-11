package main

import "fmt"

// maps: think dicts in py or objects, more or less, in js

func main() {
	// made up of key-value pairs
	// keys, in different maps,  can be of different types
	// keys, in a map, must be of the same type. case applies for values
	//Undordered (as of Jan 11, 2022)

	//strs as keys: example; menu
	menu := map[string]float64{
		"Ugali":                 20.00,
		"Chips":                 100.00,
		"Mutton | Goat":         240.00,
		"Fish":                  250.00,
		"Chicken":               200.00,
		"Ugali & Mutton | Goat": 250.00,
		"Fish & Chips":          300.00,
		"Chicken & Chips":       270.00,
	}

	fmt.Println("Entire menu as a map object:")
	fmt.Println(menu)
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print out one menu item...")
	fmt.Println("Price of Ugali is:", menu["Ugali"])
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Traverse the map:")
	for k, v := range menu {
		fmt.Printf("The price of %v is %0.2f\n", k, v)
	}

	//ints as keys: example; phonebook
	phonebook := map[int]string{
		254712345678: "Goat Matata",
		254723456789: "Dada Ng'ombe",
		254734567890: "Kaka Mbweha",
		254745678901: "Mzee Kobe",
		254756789012: "Kaka Mbwamwitu",
		254767890123: "Kaka Dubu",
	}

	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print entire phonebook as a map obj")
	fmt.Println(phonebook)
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print out one number:")
	fmt.Println("The number 254734567890 belongs to:", phonebook[254734567890])
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Traverse the map:")
	for k, v := range phonebook {
		fmt.Printf("%v's number is %v\n", v, k)
	}

	//Updating a map
	menu["Ugali"] = 40
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Edited menu as a map object:")
	fmt.Println(menu)
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print out edited menu item...")
	fmt.Println("Price of Ugali is:", menu["Ugali"])
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Traverse the mutated map:")
	for k, v := range menu {
		fmt.Printf("The price of %v is %0.2f\n", k, v)
	}

	//adding a k:v pair to a map
	phonebook[254778901234] = "Paka the Cat"

	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print edited phonebook as a map obj")
	fmt.Println(phonebook)
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Print out one number:")
	fmt.Println("The number 254734567890 belongs to:", phonebook[254734567890])
	fmt.Println("*-*-*-*-*-*-*")
	fmt.Println("Traverse the mutated map:")
	for k, v := range phonebook {
		fmt.Printf("%v's number is %v\n", v, k)
	}

}
