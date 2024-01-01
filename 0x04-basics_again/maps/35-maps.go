package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok:=users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil
}

type user struct {
	name string
	number int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("=================================================")

	deleted, err := deleteIfNecessary(users, name)
	if err!=nil {
		fmt.Println(err)
		return
	}

	if deleted {
		fmt.Println("Deleted ", name)
		return
	}

	fmt.Println("Did not delete", name)
}

func main()  {
	users := map[string]user{
		"Goat Matata": {
			name: "Goat Matata",
			number: 254701234567,
			scheduledForDeletion: false,
		},
		"Dada Ng'ombe": {
			name: "Dada Ng'ombe",
			number: 254701234567,
			scheduledForDeletion: true,
		},
		"Kaka Dubu": {
			name: "Kaka Dubu",
			number: 254701234567,
			scheduledForDeletion: false,
		},
		"Mzee Kobe": {
			name: "Mzee Kobe",
			number: 254701234567,
			scheduledForDeletion: true,
		},
		"Bi Buibui": {
			name: "Bi Buibui",
			number: 254701234567,
			scheduledForDeletion: true,
		},
		"Kaka Tai": {
			name: "Kaka Tai",
			number: 254701234567,
			scheduledForDeletion: false,
		},
		"Kaka Mbwamwitu": {
			name: "Kaka Mbwamwitu",
			number: 254701234567,
			scheduledForDeletion: false,
		},
	}

	test(users, "Goat Matata")
	test(users, "Kaka Mbweha")
	test(users, "Bi Buibui")
}