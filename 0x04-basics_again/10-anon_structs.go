package main

import "fmt"

myCar := struct {
	Make string
	Model string
} {
	Make: "datsun",
	Model: "240z",
}

type car struct{
	Make string
	Model string
	Height int
	Width int
	Wheel struct {
		Radius int
		Material string
	}
}

func main()  {
	fmt.Println("myCar", myCar)

	anotherCar := car{
		Make: "ferrari",
		Model: "250 gt",
		Height: 155,
		Width: 195,
		Wheel: {
			Radius: 125,
			Material: "rubber",
		},
	}
	fmt.Println("anotherCar", anotherCar)
}