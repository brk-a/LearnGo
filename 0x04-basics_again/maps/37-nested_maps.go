package main

import "fmt"


func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if name=="" {
			continue
		}

		firstChar := name[0]
		_, ok := counts[rune(firstChar)]
		if !ok {
			counts[rune(firstChar)] = make(map[string]int)
		}
		counts[rune(firstChar)][name]++
	}

	return counts
}

func test(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("=====================================")
}

func main() {
	test(getNames(5), 'M', "Mzee Kobe")
	test(getNames(10), 'G', "Goat Matata")
	test(getNames(15), 'D', "Kaka Dubu")
	test(getNames(20), 'P', "Paka the Cat")
	test(getNames(25), 'F', "F Njakai")
	test(getNames(30), 'M', "Mzee Kobe")
}

func getNames(length int) []string {
	names := []string{
		"Goat Matata",
		"Kaka Tai",
		"Bi Buibui",
		"Mzee Kobe",
		"Mzee Kobe",
		"Mzee Kobe",
		"Bi Buibui",
		"Bi Buibui",
		"Paka the Cat",
		"Paka the Cat",
		"Paka the Cat",
		"F Njakai", 
		"Kaka Mbweha",
		"Kaka Dubu",
		"Dada Ng'ombe",
		"",
		"Dada Ng'ombe",
		"F Njakai",
		"Dada Ng'ombe",
		"Dada Ng'ombe",
		"Kaka Mbweha",
		"Kaka Mbwambwitu",
		"Kaka Mbwambwitu",
		"Kaka Dubu",
		"Kaka Mbwambwitu",
		"Paka the Cat",
		"Paka the Cat",
		"Goat Matata",
		"Kaka Tai",
		"Mzee Kobe",
		"",
	}
	if length > len(names) {
		length = len(names)
	}
	return names[:length]
}