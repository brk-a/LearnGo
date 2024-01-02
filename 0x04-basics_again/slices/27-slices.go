package main

import "fmt"

const (
	retry1 = "click here to sign up"
	retry2 = "please, please click here"
	retry3 = "we beg you to sign up"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

func testSend(name string, doneAt int)  {
	fmt.Printf("Sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i:=0; i<len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`Sending: "%v" `, msg)
		fmt.Println()
		if i==doneAt {
			fmt.Println("user has responded")
			break
		}
		if i==len(messages)-1 {
			fmt.Println("we failed")
		}
	}
}


func main() {
	testSend("Goat Matata", 0)
	testSend("Dada Ng'ombe", 1)
	testSend("Kaka Dubu", 2)
	testSend("Mzee Kobe", 3)
}