package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, world!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString=="" {
		log.Fatal("port could not be found in env")
	}

	fmt.Println("Port: ", portString)
}
