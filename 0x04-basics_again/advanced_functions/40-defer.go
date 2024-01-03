package main

import (
	"fmt"
	"sort"
)

const (
	logDeleted = "user deleted"
	logNotFound = "user not found"
	logAdmin = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	
	user, ok :=  users[name]
	if !ok {
		// delete(users, name)
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	// delete(users, name)
	return logDeleted
}

type user struct {
	name string
	number int
	admin bool
}

func test(users map[string]user, name string) {
	fmt.Printf("attempting to delete %s...\n", name)
	defer fmt.Println("======================================================")

	log := logAndDelete(users, name)
	fmt.Println("Log: ", log)
}

func main() {
	users := map[string]user{
		"Goat Matata": {
			name: "Goat Matata",
			number: 254701234567,
			admin: true,
		},
		"Dada Ng'ombe": {
			name: "Dada Ng'ombe",
			number: 254712345678,
			admin: false,
		},
		"Kaka Mbweha": {
			name: "Kaka Mbweha",
			number: 254723456789,
			admin: false,
		},
		"Kaka Dubu": {
			name: "Kaka Dubu",
			number: 254734567890,
			admin: false,
		},
		"Kaka Mbwamwitu": {
			name: "Kaka Mbwamwitu",
			number: 254745678901,
			admin: true,
		},
	}

	fmt.Println("initial user: ")
	usersSorted := []string{}

	for name:= range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)

	for _, name:= range usersSorted {
		fmt.Println("- ", name)
	}
	fmt.Println("=====================================")

	test(users, "Goat Matata")
	test(users, "Paka the Cat")
	test(users, "Kaka Mbweha")

	fmt.Println("final users: ")
	usersSorted = []string{}

	for name:=range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)

	for _, name:=range usersSorted {
		fmt.Println("- ", name)
	}
	fmt.Println("============================================")
}