package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error){
	if len(names)!=len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userMap := make(map[string]user)
	for i:=0; i<len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user {
			name: name,
			phoneNumber: phoneNumber,
		}
	}
	
	/**
	alternative
	for _, name := range names {
		for _, phoneNumber := range phoneNumbers {
			userMap[name] = user {
				name: name,
				phoneNumber: phoneNumber,
			}
		}
	}
	*/

	return userMap, nil
}

type user struct {
	name string
	phoneNumber int
}

func test(names []string, phoneNumbers []int)  {
	fmt.Println("Creating map...")
	defer fmt.Println("========================================================")

	users, err := getUserMap(names, phoneNumbers)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println("- name: ", users[name].name)
		fmt.Println("- phone number: ", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"Goat Matata", "Dada Ng'ombe", "Mzee Kobe"},
		[]int{254701234567, 254712345678, 254723456789},
	)
	test(
		[]string{"Kaka Mbweha", "Kaka Tai"},
		[]int{254734567890, 254745678901, 254756789012},
	)
	test(
		[]string{"Bi Buibui", "Paka the Cat", "Kaka Mbwamwitu", "Kaka Dubu"},
		[]int{254767890123, 254778901234, 254789012345, 254790123456},
	)
}